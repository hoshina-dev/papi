package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	appConfig "github.com/hoshina-dev/papi/internal/config"
	"github.com/hoshina-dev/papi/internal/graphql"
	"github.com/hoshina-dev/papi/internal/handler"
	"github.com/hoshina-dev/papi/internal/infra/postgres"
	"github.com/hoshina-dev/papi/internal/infra/rabbitmq"
	storage "github.com/hoshina-dev/papi/internal/infra/s3"
	"github.com/hoshina-dev/papi/internal/repository"
	"github.com/hoshina-dev/papi/internal/server"
	"github.com/hoshina-dev/papi/internal/service"
)

// App holds the running components of the application.
// Build wires everything; Run starts the goroutines.
type App struct {
	fiber        *fiber.App
	rabbitKeeper *rabbitmq.ResilientPublisher
	port         string
}

// Build constructs every dependency and returns a ready-to-run App.
//
// Failures are treated by severity:
//   - Postgres: fatal — the app is useless without a database.
//   - S3 client: fatal — presigned URLs are core to the API.
//   - RabbitMQ: non-fatal — the keeper goroutine will connect once Run starts;
//     optimization mutations will return errors until the connection is up.
func Build(cfg *appConfig.Config) (*App, error) {
	// --- Postgres ---
	db, err := postgres.Connect(cfg.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	// --- S3 / Cloudflare R2 ---
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(cfg.S3Region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("aws config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.S3Endpoint)
		o.UsePathStyle = true
	})
	s3Storage := storage.NewS3StorageService(s3Client, cfg.S3Bucket, cfg.S3BaseURL)

	// --- RabbitMQ ---
	// Does not connect here; Run() starts the keeper goroutine.
	rabbitKeeper := rabbitmq.NewResilientPublisher(cfg.RabbitMQURL, cfg.RabbitMQExchange)

	// --- Repositories ---
	partRepo := repository.NewPartRepository(db)
	manufacturerRepo := repository.NewManufacturerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	partsInventoryRepo := repository.NewPartsInventoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	productPartRepo := repository.NewProductPartRepository(db)
	productInventoryRepo := repository.NewProductInventoryRepository(db)
	part3DModelRepo := repository.NewModel3DRepository(db)
	jobLogRepo := repository.NewOptimizationJobLogRepository(db)

	// --- Services ---
	webhookURL := cfg.OptimizationWebhookURL
	if webhookURL == "" {
		webhookURL = fmt.Sprintf("http://localhost:%s/webhook/optimization", cfg.Port)
		log.Printf("[app] OPTIMIZATION_WEBHOOK_URL not set, defaulting to: %s", webhookURL)
	}

	partSvc := service.NewPartService(partRepo, manufacturerRepo, categoryRepo)
	manufacturerSvc := service.NewManufacturerService(manufacturerRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	partsInventorySvc := service.NewPartsInventoryService(partsInventoryRepo, partRepo)
	productSvc := service.NewProductService(productRepo, partRepo, productPartRepo)
	productInventorySvc := service.NewProductInventoryService(productInventoryRepo, productRepo, partsInventoryRepo)
	storageSvc := service.NewStorageService(s3Storage)
	optimizationSvc := service.NewOptimizationService(
		s3Storage,
		rabbitKeeper,
		part3DModelRepo,
		webhookURL,
		cfg.RabbitMQExchange,
		cfg.RabbitMQRoutingKey,
	)

	// --- HTTP layer ---
	webhookHandler := handler.NewWebhookHandler(part3DModelRepo, jobLogRepo)
	resolver := graphql.NewResolver(
		partSvc, manufacturerSvc, categorySvc,
		partsInventorySvc, productSvc, productInventorySvc,
		storageSvc, optimizationSvc,
	)

	// Health check closure — captures db and rabbitKeeper.
	// Postgres: critical (503 if down). RabbitMQ: non-critical (degraded).
	// R2 is intentionally excluded — no reliable ping endpoint on Cloudflare R2.
	healthHandler := func(c *fiber.Ctx) error {
		components := fiber.Map{}
		overall := "ok"

		sqlDB, err := db.DB()
		if err != nil || sqlDB.PingContext(c.UserContext()) != nil {
			components["postgres"] = "error"
			overall = "error"
		} else {
			components["postgres"] = "ok"
		}

		if rabbitKeeper.IsConnected() {
			components["rabbitmq"] = "ok"
		} else {
			components["rabbitmq"] = "unavailable"
			if overall == "ok" {
				overall = "degraded"
			}
		}

		status := http.StatusOK
		if overall == "error" {
			status = http.StatusServiceUnavailable
		}

		return c.Status(status).JSON(fiber.Map{
			"status":     overall,
			"components": components,
		})
	}

	fiberApp := server.New(resolver, webhookHandler, healthHandler, cfg.CORSOrigins)

	return &App{
		fiber:        fiberApp,
		rabbitKeeper: rabbitKeeper,
		port:         cfg.Port,
	}, nil
}

// Run starts all background goroutines and blocks until they all exit.
//
// Goroutine responsibilities:
//   - RabbitMQ keeper: reconnects silently in the background; a failure here
//     never tears down the HTTP server.
//   - HTTP server: if Listen returns an error the whole process should restart,
//     so the error is propagated through the errgroup.
//
// Cancel ctx (e.g. via OS signal) to trigger a graceful shutdown of both.
func (a *App) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// Goroutine 1: RabbitMQ connection keeper.
	// Errors are logged internally and retried — never propagated upward.
	g.Go(func() error {
		a.rabbitKeeper.Run(ctx)
		return nil
	})

	// Goroutine 2: HTTP server (GraphQL + webhook).
	// A startup or runtime error here is fatal for the process.
	g.Go(func() error {
		return runHTTPServer(ctx, a.fiber, a.port)
	})

	return g.Wait()
}

func runHTTPServer(ctx context.Context, app *fiber.App, port string) error {
	listenErr := make(chan error, 1)
	go func() {
		listenErr <- app.Listen(":" + port)
	}()

	select {
	case err := <-listenErr:
		return err
	case <-ctx.Done():
		return app.ShutdownWithTimeout(10 * time.Second)
	}
}
