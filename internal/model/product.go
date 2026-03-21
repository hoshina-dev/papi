package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string
	Version     *string
	Description *string
	Parts       []ProductPart `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}

func (Product) TableName() string {
	return "products"
}
