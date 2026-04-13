package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string
	Version     *string
	Description *string
	Images      pq.StringArray `gorm:"type:text[];default:'{}'"`
	Parts       []ProductPart  `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}

func (Product) TableName() string {
	return "products"
}
