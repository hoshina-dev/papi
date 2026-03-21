package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductInventory struct {
	ID           uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProductID    uuid.UUID        `gorm:"type:uuid;not null"`
	Product      Product
	SerialNumber string
	IsAvailable  bool             `gorm:"default:true"`
	Notes        *string
	PartsUsed    []PartsInventory `gorm:"many2many:product_inventory_parts;joinForeignKey:ProductInventoryId;joinReferences:PartInventoryId"`
	CreatedAt    time.Time        `gorm:"autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}

func (ProductInventory) TableName() string {
	return "product_inventory"
}
