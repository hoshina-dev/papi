package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"github.com/hoshina-dev/papi/internal/repository"
)

type PartsInventoryService struct {
	repo     repository.PartsInventoryRepository
	partRepo repository.PartRepository
}

func NewPartsInventoryService(repo repository.PartsInventoryRepository, partRepo repository.PartRepository) *PartsInventoryService {
	return &PartsInventoryService{repo: repo, partRepo: partRepo}
}

func (s *PartsInventoryService) GetAll(ctx context.Context) ([]model.PartsInventory, error) {
	return s.repo.GetAll(ctx)
}

func (s *PartsInventoryService) GetByID(ctx context.Context, id uuid.UUID) (*model.PartsInventory, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PartsInventoryService) GetByPartID(ctx context.Context, partID uuid.UUID) ([]model.PartsInventory, error) {
	return s.repo.GetByPartID(ctx, partID)
}

func (s *PartsInventoryService) Create(ctx context.Context, input model.CreatePartsInventoryInput) (*model.PartsInventory, error) {
	part, err := s.partRepo.GetByID(ctx, input.PartID)
	if err != nil {
		return nil, fmt.Errorf("part with id %s not found", input.PartID)
	}

	isAvailable := true
	if input.IsAvailable != nil {
		isAvailable = *input.IsAvailable
	}

	item := &model.PartsInventory{
		PartID:       input.PartID,
		Part:         *part,
		SerialNumber: input.SerialNumber,
		IsAvailable:  isAvailable,
		Notes:        input.Notes,
	}
	if err := s.repo.Create(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *PartsInventoryService) Update(ctx context.Context, id uuid.UUID, input model.UpdatePartsInventoryInput) (*model.PartsInventory, error) {
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

func (s *PartsInventoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
