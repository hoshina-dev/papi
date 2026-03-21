package model

import (
	"github.com/google/uuid"
	"github.com/hoshina-dev/papi/internal/scalar"
)

// ── Part ─────────────────────────────────────────────────────────────────────

type CreatePartInput struct {
	Name             string    `validate:"required"`
	PartNumber       string    `validate:"required"`
	ManufacturerID   uuid.UUID `validate:"required,uuid4"`
	Description      *string
	TemperatureStage *string
	Specifications   scalar.JSON
	Images           []string    `validate:"omitempty,dive,url"`
	CategoryIDs      []uuid.UUID `validate:"omitempty,dive,uuid4"`
}

type UpdatePartInput struct {
	Name             *string
	Description      *string
	TemperatureStage *string
	Specifications   scalar.JSON
	Images           []string    `validate:"omitempty,dive,url"`
	CategoryIDs      []uuid.UUID `validate:"omitempty,dive,uuid4"`
}

// ── Category ──────────────────────────────────────────────────────────────────

type CreateCategoryInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateCategoryInput struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// ── Manufacturer ──────────────────────────────────────────────────────────────

type CreateManufacturerInput struct {
	Name            string  `json:"name"`
	CountryOfOrigin *string `json:"countryOfOrigin,omitempty"`
}

type UpdateManufacturerInput struct {
	Name            *string `json:"name,omitempty"`
	CountryOfOrigin *string `json:"countryOfOrigin,omitempty"`
}

// ── PartsInventory ────────────────────────────────────────────────────────────

type CreatePartsInventoryInput struct {
	PartID       uuid.UUID `validate:"required,uuid4"`
	SerialNumber string    `validate:"required"`
	IsAvailable  *bool
	Notes        *string
}

type UpdatePartsInventoryInput struct {
	SerialNumber *string
	IsAvailable  *bool
	Notes        *string
}

// ── Product ───────────────────────────────────────────────────────────────────

type CreateProductInput struct {
	Name        string `validate:"required"`
	Version     *string
	Description *string
}

type UpdateProductInput struct {
	Name        *string
	Version     *string
	Description *string
}

// ── ProductPart (BOM) ─────────────────────────────────────────────────────────

type AddProductPartInput struct {
	ProductID uuid.UUID `validate:"required,uuid4"`
	PartID    uuid.UUID `validate:"required,uuid4"`
	Quantity  int32     `validate:"required,gt=0"`
	Notes     *string
}

type UpdateProductPartInput struct {
	Quantity *int32 `validate:"omitempty,gt=0"`
	Notes    *string
}

// ── ProductInventory ──────────────────────────────────────────────────────────

type CreateProductInventoryInput struct {
	ProductID    uuid.UUID `validate:"required,uuid4"`
	SerialNumber string    `validate:"required"`
	IsAvailable  *bool
	Notes        *string
}

type UpdateProductInventoryInput struct {
	SerialNumber *string
	IsAvailable  *bool
	Notes        *string
}
