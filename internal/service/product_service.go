package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"github.com/hoshina-dev/papi/internal/repository"
)

type ProductService struct {
	repo            repository.ProductRepository
	partRepo        repository.PartRepository
	productPartRepo repository.ProductPartRepository
}

func NewProductService(repo repository.ProductRepository, partRepo repository.PartRepository, productPartRepo repository.ProductPartRepository) *ProductService {
	return &ProductService{repo: repo, partRepo: partRepo, productPartRepo: productPartRepo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]model.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Search(ctx context.Context, name string) ([]model.Product, error) {
	return s.repo.Search(ctx, name)
}

func (s *ProductService) Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	p := &model.Product{
		Name:        input.Name,
		Version:     input.Version,
		Description: input.Description,
	}
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Update(ctx context.Context, id uuid.UUID, input model.UpdateProductInput) (*model.Product, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		p.Name = *input.Name
	}
	if input.Version != nil {
		p.Version = input.Version
	}
	if input.Description != nil {
		p.Description = input.Description
	}

	if err := s.repo.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) AddPart(ctx context.Context, input model.AddProductPartInput) (*model.ProductPart, error) {
	_, err := s.repo.GetByID(ctx, input.ProductID)
	if err != nil {
		return nil, fmt.Errorf("product with id %s not found", input.ProductID)
	}

	part, err := s.partRepo.GetByID(ctx, input.PartID)
	if err != nil {
		return nil, fmt.Errorf("part with id %s not found", input.PartID)
	}

	pp := &model.ProductPart{
		ProductID: input.ProductID,
		PartID:    input.PartID,
		Part:      *part,
		Quantity:  input.Quantity,
		Notes:     input.Notes,
	}
	if err := s.productPartRepo.Create(ctx, pp); err != nil {
		return nil, err
	}
	return pp, nil
}

func (s *ProductService) UpdatePart(ctx context.Context, id uuid.UUID, input model.UpdateProductPartInput) (*model.ProductPart, error) {
	pp, err := s.productPartRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Quantity != nil {
		pp.Quantity = *input.Quantity
	}
	if input.Notes != nil {
		pp.Notes = input.Notes
	}

	if err := s.productPartRepo.Update(ctx, pp); err != nil {
		return nil, err
	}
	return pp, nil
}

func (s *ProductService) RemovePart(ctx context.Context, id uuid.UUID) error {
	return s.productPartRepo.Delete(ctx, id)
}
