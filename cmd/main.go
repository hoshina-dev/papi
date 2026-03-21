package main

import (
	"log"

	appConfig "github.com/hoshina-dev/pasta/internal/config"
	"github.com/hoshina-dev/pasta/internal/graphql"
	"github.com/hoshina-dev/pasta/internal/infra/postgres"
	"github.com/hoshina-dev/pasta/internal/repository"
	"github.com/hoshina-dev/pasta/internal/server"
	"github.com/hoshina-dev/pasta/internal/service"
)

func main() {
	cfg := appConfig.Load()

	db, err := postgres.Connect(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	partRepo := repository.NewPartRepository(db)
	manufacturerRepo := repository.NewManufacturerRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	partsInventoryRepo := repository.NewPartsInventoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	productPartRepo := repository.NewProductPartRepository(db)
	productInventoryRepo := repository.NewProductInventoryRepository(db)

	partSvc := service.NewPartService(partRepo, manufacturerRepo, categoryRepo)
	manufacturerSvc := service.NewManufacturerService(manufacturerRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	partsInventorySvc := service.NewPartsInventoryService(partsInventoryRepo, partRepo)
	productSvc := service.NewProductService(productRepo, partRepo, productPartRepo)
	productInventorySvc := service.NewProductInventoryService(productInventoryRepo, productRepo, partsInventoryRepo)

	resolver := graphql.NewResolver(partSvc, manufacturerSvc, categorySvc, partsInventorySvc, productSvc, productInventorySvc)

	app := server.New(resolver, cfg.CORSOrigins)

	log.Printf("starting server on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
