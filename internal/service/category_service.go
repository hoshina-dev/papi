package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/hoshina-dev/pasta/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	c := &model.Category{
		Name:        input.Name,
		Description: input.Description,
	}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) Update(ctx context.Context, id uuid.UUID, input model.UpdateCategoryInput) (*model.Category, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, fmt.Errorf("category with id %s not found", id)
	}
	model.ApplyUpdateCategoryInput(c, input)
	if err := s.repo.Update(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CategoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
