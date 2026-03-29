package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"github.com/hoshina-dev/papi/internal/repository"
)

type ProductInventoryService struct {
	repo               repository.ProductInventoryRepository
	productRepo        repository.ProductRepository
	partsInventoryRepo repository.PartsInventoryRepository
}

func NewProductInventoryService(repo repository.ProductInventoryRepository, productRepo repository.ProductRepository, partsInventoryRepo repository.PartsInventoryRepository) *ProductInventoryService {
	return &ProductInventoryService{repo: repo, productRepo: productRepo, partsInventoryRepo: partsInventoryRepo}
}

func (s *ProductInventoryService) GetAll(ctx context.Context) ([]model.ProductInventory, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductInventoryService) GetByID(ctx context.Context, id uuid.UUID) (*model.ProductInventory, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductInventoryService) GetByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductInventory, error) {
	return s.repo.GetByProductID(ctx, productID)
}

func (s *ProductInventoryService) Create(ctx context.Context, input model.CreateProductInventoryInput) (*model.ProductInventory, error) {
	product, err := s.productRepo.GetByID(ctx, input.ProductID)
	if err != nil {
		return nil, fmt.Errorf("product with id %s not found", input.ProductID)
	}

	isAvailable := true
	if input.IsAvailable != nil {
		isAvailable = *input.IsAvailable
	}

	item := &model.ProductInventory{
		ProductID:    input.ProductID,
		Product:      *product,
		SerialNumber: input.SerialNumber,
		IsAvailable:  isAvailable,
		Notes:        input.Notes,
	}
	if err := s.repo.Create(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ProductInventoryService) Update(ctx context.Context, id uuid.UUID, input model.UpdateProductInventoryInput) (*model.ProductInventory, error) {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.SerialNumber != nil {
		item.SerialNumber = *input.SerialNumber
	}
	if input.IsAvailable != nil {
		item.IsAvailable = *input.IsAvailable
	}
	if input.Notes != nil {
		item.Notes = input.Notes
	}

	if err := s.repo.Update(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ProductInventoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductInventoryService) AddPartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	_, err := s.repo.GetByID(ctx, productInventoryID)
	if err != nil {
		return fmt.Errorf("product inventory item with id %s not found", productInventoryID)
	}

	_, err = s.partsInventoryRepo.GetByID(ctx, partsInventoryID)
	if err != nil {
		return fmt.Errorf("parts inventory item with id %s not found", partsInventoryID)
	}

	return s.repo.AddPartsInventory(ctx, productInventoryID, partsInventoryID)
}

func (s *ProductInventoryService) RemovePartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	return s.repo.RemovePartsInventory(ctx, productInventoryID, partsInventoryID)
}
