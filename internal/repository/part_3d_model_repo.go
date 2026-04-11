package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/model"
	"gorm.io/gorm"
)

type Part3DModelRepository interface {
	Create(ctx context.Context, m *model.Model3D) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Model3D, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status model.Model3DStatus) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type part3DModelRepository struct {
	db *gorm.DB
}

func NewPart3DModelRepository(db *gorm.DB) Part3DModelRepository {
	return &part3DModelRepository{db: db}
}

func (r *part3DModelRepository) Create(ctx context.Context, m *model.Model3D) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *part3DModelRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Model3D, error) {
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

func (r *part3DModelRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status model.Model3DStatus) error {
	return r.db.WithContext(ctx).Model(&model.Model3D{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status}).Error
}

func (r *part3DModelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	res := r.db.WithContext(ctx).Delete(&model.Model3D{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("part 3d model not found")
	}
	return nil
}
