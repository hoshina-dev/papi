package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error)
	Search(ctx context.Context, name string) ([]model.Product, error)
	Create(ctx context.Context, p *model.Product) error
	Update(ctx context.Context, p *model.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductPartRepository interface {
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductPart, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.ProductPart, error)
	Create(ctx context.Context, pp *model.ProductPart) error
	Update(ctx context.Context, pp *model.ProductPart) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// ── Product ───────────────────────────────────────────────────────────────────

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	var items []model.Product
	err := r.db.WithContext(ctx).Preload("Parts.Part").Find(&items).Error
	return items, err
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	var p model.Product
	err := r.db.WithContext(ctx).Preload("Parts.Part").First(&p, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Search(ctx context.Context, name string) ([]model.Product, error) {
	var items []model.Product
	name = strings.ReplaceAll(name, `\`, `\\`)
	name = strings.ReplaceAll(name, `%`, `\%`)
	name = strings.ReplaceAll(name, `_`, `\_`)
	err := r.db.WithContext(ctx).Preload("Parts.Part").Where("name ILIKE ? ESCAPE '\\'", "%"+name+"%").Find(&items).Error
	return items, err
}

func (r *productRepository) Create(ctx context.Context, p *model.Product) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *productRepository) Update(ctx context.Context, p *model.Product) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(p).Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Product{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

// ── ProductPart ───────────────────────────────────────────────────────────────

type productPartRepository struct {
	db *gorm.DB
}

func NewProductPartRepository(db *gorm.DB) ProductPartRepository {
	return &productPartRepository{db: db}
}

func (r *productPartRepository) GetByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductPart, error) {
	var items []model.ProductPart
	err := r.db.WithContext(ctx).Preload("Part").Where("product_id = ?", productID).Find(&items).Error
	return items, err
}

func (r *productPartRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.ProductPart, error) {
	var pp model.ProductPart
	err := r.db.WithContext(ctx).Preload("Part").First(&pp, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pp, nil
}

func (r *productPartRepository) Create(ctx context.Context, pp *model.ProductPart) error {
	return r.db.WithContext(ctx).Create(pp).Error
}

func (r *productPartRepository) Update(ctx context.Context, pp *model.ProductPart) error {
	return r.db.WithContext(ctx).Clauses(clause.Returning{}).Save(pp).Error
}

func (r *productPartRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.ProductPart{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("product part not found")
	}
	return nil
}
