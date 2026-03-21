package graphql

import "github.com/hoshina-dev/pasta/internal/service"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	partService             *service.PartService
	manufacturerService     *service.ManufacturerService
	categoryService         *service.CategoryService
	partsInventoryService   *service.PartsInventoryService
	productService          *service.ProductService
	productInventoryService *service.ProductInventoryService
}

func NewResolver(
	partSvc *service.PartService,
	manufacturerSvc *service.ManufacturerService,
	categorySvc *service.CategoryService,
	partsInventorySvc *service.PartsInventoryService,
	productSvc *service.ProductService,
	productInventorySvc *service.ProductInventoryService,
) *Resolver {
	return &Resolver{
		partService:             partSvc,
		manufacturerService:     manufacturerSvc,
		categoryService:         categorySvc,
		partsInventoryService:   partsInventorySvc,
		productService:          productSvc,
		productInventoryService: productInventorySvc,
	}
}
