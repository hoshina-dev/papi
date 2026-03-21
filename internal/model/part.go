package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/scalar"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Part struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name             string
	PartNumber       string
	ManufacturerID   uuid.UUID
	Manufacturer     Manufacturer
	Description      *string
	TemperatureStage *string
	Specifications   scalar.JSON    `gorm:"type:jsonb"`
	Images           pq.StringArray `gorm:"type:text[];default:'{}'"`
	Categories       []Category     `gorm:"many2many:part_categories;joinForeignKey:PartId"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt
}

func (Part) TableName() string {
	return "parts"
}
