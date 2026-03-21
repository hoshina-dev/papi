package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoshina-dev/papi/internal/app"
	appConfig "github.com/hoshina-dev/papi/internal/config"
)

func main() {
	cfg := appConfig.Load()

	a, err := app.Build(cfg)
	if err != nil {
		log.Fatalf("failed to build app: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("starting server on :%s", cfg.Port)
	// running goroutines for: RabbitMQ connection keeper, HTTP server (graphql + webhook callback handler)
	if err := a.Run(ctx); err != nil {
		log.Fatalf("app exited with error: %v", err)
	}
}
