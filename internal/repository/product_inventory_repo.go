package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductInventoryRepository interface {
	GetAll(ctx context.Context) ([]model.ProductInventory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.ProductInventory, error)
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductInventory, error)
	Create(ctx context.Context, item *model.ProductInventory) error
	Update(ctx context.Context, item *model.ProductInventory) error
	Delete(ctx context.Context, id uuid.UUID) error
	AddPartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error
	RemovePartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error
}

type productInventoryRepository struct {
	db *gorm.DB
}

func NewProductInventoryRepository(db *gorm.DB) ProductInventoryRepository {
	return &productInventoryRepository{db: db}
}

func (r *productInventoryRepository) GetAll(ctx context.Context) ([]model.ProductInventory, error) {
	var items []model.ProductInventory
	err := r.db.WithContext(ctx).Preload("Product").Preload("PartsUsed.Part").Find(&items).Error
	return items, err
}

func (r *productInventoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.ProductInventory, error) {
	var item model.ProductInventory
	err := r.db.WithContext(ctx).Preload("Product").Preload("PartsUsed.Part").First(&item, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *productInventoryRepository) GetByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductInventory, error) {
	var items []model.ProductInventory
	err := r.db.WithContext(ctx).Preload("Product").Preload("PartsUsed.Part").Where("product_id = ?", productID).Find(&items).Error
	return items, err
}

func (r *productInventoryRepository) Create(ctx context.Context, item *model.ProductInventory) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *productInventoryRepository) Update(ctx context.Context, item *model.ProductInventory) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(item).Error
}

func (r *productInventoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.ProductInventory{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("product inventory item not found")
	}
	return nil
}

func (r *productInventoryRepository) AddPartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	partsItem := model.PartsInventory{ID: partsInventoryID}
	return r.db.WithContext(ctx).
		Model(&model.ProductInventory{ID: productInventoryID}).
		Association("PartsUsed").
		Append(&partsItem)
}

func (r *productInventoryRepository) RemovePartsInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	partsItem := model.PartsInventory{ID: partsInventoryID}
	return r.db.WithContext(ctx).
		Model(&model.ProductInventory{ID: productInventoryID}).
		Association("PartsUsed").
		Delete(&partsItem)
}
