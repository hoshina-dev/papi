package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Part3DModelStatus string

const (
	Part3DModelStatusProcessing Part3DModelStatus = "processing"
	Part3DModelStatusReady      Part3DModelStatus = "ready"
	Part3DModelStatusFailed     Part3DModelStatus = "failed"
)

func (Part3DModel) TableName() string { return "part_3d_models" }

type Part3DModel struct {
	ID           uuid.UUID         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RawURL       string            `gorm:"type:text;not null"`
	ProcessedKey *string           `gorm:"type:text"`
	Status       Part3DModelStatus `gorm:"type:text;not null;default:'processing'"`
	CreatedAt    time.Time         `gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdatedAt    time.Time         `gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt    `gorm:"type:timestamp with time zone"`
}
