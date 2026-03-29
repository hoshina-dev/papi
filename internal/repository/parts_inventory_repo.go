package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PartsInventoryRepository interface {
	GetAll(ctx context.Context) ([]model.PartsInventory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.PartsInventory, error)
	GetByPartID(ctx context.Context, partID uuid.UUID) ([]model.PartsInventory, error)
	Create(ctx context.Context, item *model.PartsInventory) error
	Update(ctx context.Context, item *model.PartsInventory) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type partsInventoryRepository struct {
	db *gorm.DB
}

func NewPartsInventoryRepository(db *gorm.DB) PartsInventoryRepository {
	return &partsInventoryRepository{db: db}
}

func (r *partsInventoryRepository) GetAll(ctx context.Context) ([]model.PartsInventory, error) {
	var items []model.PartsInventory
	err := r.db.WithContext(ctx).Preload("Part").Find(&items).Error
	return items, err
}

func (r *partsInventoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.PartsInventory, error) {
	var item model.PartsInventory
	err := r.db.WithContext(ctx).Preload("Part").First(&item, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *partsInventoryRepository) GetByPartID(ctx context.Context, partID uuid.UUID) ([]model.PartsInventory, error) {
	var items []model.PartsInventory
	err := r.db.WithContext(ctx).Preload("Part").Where("part_id = ?", partID).Find(&items).Error
	return items, err
}

func (r *partsInventoryRepository) Create(ctx context.Context, item *model.PartsInventory) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *partsInventoryRepository) Update(ctx context.Context, item *model.PartsInventory) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(item).Error
}

func (r *partsInventoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.PartsInventory{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("parts inventory item not found")
	}
	return nil
}
