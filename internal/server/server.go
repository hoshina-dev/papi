package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hoshina-dev/papi/internal/graphql"
	 webhookHandler "github.com/hoshina-dev/papi/internal/handler"
)

func New(resolver *graphql.Resolver, webhookHandler *webhookHandler.WebhookHandler, healthHandler fiber.Handler, corsOrigins string) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: corsOrigins,
	}))

	srv := handler.New(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: resolver,
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	app.All("/graphql", adaptor.HTTPHandler(srv))
	app.Get("/", adaptor.HTTPHandler(
		playground.Handler("papi GraphQL", "/graphql"),
	))
	app.Get("/health", healthHandler)

	app.Post("/webhook/optimization", webhookHandler.HandleOptimizationCallback)

	return app
}
