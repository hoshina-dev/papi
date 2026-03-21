package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PartsInventory struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PartID       uuid.UUID `gorm:"type:uuid;not null"`
	Part         Part
	SerialNumber string
	IsAvailable  bool      `gorm:"default:true"`
	Notes        *string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}

func (PartsInventory) TableName() string {
	return "parts_inventory"
}
