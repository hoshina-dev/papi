package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model3DStatus string

const (
	Model3DStatusProcessing Model3DStatus = "processing"
	Model3DStatusReady      Model3DStatus = "ready"
	Model3DStatusFailed     Model3DStatus = "failed"
)

func (Model3D) TableName() string { return "model_3d" }

type Model3D struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PartID       *uuid.UUID `gorm:"type:uuid"`
	Part         *Part
	ProductID    *uuid.UUID `gorm:"type:uuid"`
	Product      *Product
	RawKey       string         `gorm:"type:text;not null;column:raw_key"`
	ProcessedKey *string        `gorm:"type:text"`
	Status       Model3DStatus  `gorm:"type:text;not null;default:'processing'"`
	CreatedAt    time.Time      `gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp with time zone"`
}
