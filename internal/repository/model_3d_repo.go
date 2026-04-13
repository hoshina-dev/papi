package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"gorm.io/gorm"
)

type Model3DRepository interface {
	Create(ctx context.Context, m *model.Model3D) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Model3D, error)
	GetReadyByPartID(ctx context.Context, partID uuid.UUID) ([]model.Model3D, error)
	GetReadyByProductID(ctx context.Context, productID uuid.UUID) ([]model.Model3D, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status model.Model3DStatus) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type model3DRepository struct {
	db *gorm.DB
}

func NewModel3DRepository(db *gorm.DB) Model3DRepository {
	return &model3DRepository{db: db}
}

func (r *model3DRepository) Create(ctx context.Context, m *model.Model3D) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *model3DRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Model3D, error) {
	var m model.Model3D
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByPartID implements [Model3DRepository].
func (r *model3DRepository) GetReadyByPartID(ctx context.Context, partID uuid.UUID) ([]model.Model3D, error) {
	var models []model.Model3D
	err := r.db.WithContext(ctx).
		Where("part_id = ? AND status = ?", partID, model.Model3DStatusReady).
		Find(&models).Error
	return models, err
}

// GetByProductID implements [Model3DRepository].
func (r *model3DRepository) GetReadyByProductID(ctx context.Context, productID uuid.UUID) ([]model.Model3D, error) {
	var models []model.Model3D
	err := r.db.WithContext(ctx).
		Where("product_id = ? AND status = ?", productID, model.Model3DStatusReady).
		Find(&models).Error
	return models, err
}

func (r *model3DRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status model.Model3DStatus) error {
	return r.db.WithContext(ctx).Model(&model.Model3D{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status}).Error
}

func (r *model3DRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Model3D{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("part 3d model not found")
	}
	return nil
}
