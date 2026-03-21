package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductPart struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	PartID    uuid.UUID `gorm:"type:uuid;not null"`
	Part      Part
	Quantity  int32
	Notes     *string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ProductPart) TableName() string {
	return "product_parts"
}
