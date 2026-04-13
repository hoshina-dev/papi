package model

import "github.com/google/uuid"

type Optimize3DInput struct {
	PartID                    *uuid.UUID `json:"partID,omitempty"`
	ProductID                 *uuid.UUID `json:"productID,omitempty"`
	SourceURL                 string     `json:"sourceURL"`
	DracoCompressionLevel     *int32     `json:"dracoCompressionLevel,omitempty"`
	DracoPositionQuantization *int32     `json:"dracoPositionQuantization,omitempty"`
	DracoTexcoordQuantization *int32     `json:"dracoTexcoordQuantization,omitempty"`
	DracoNormalQuantization   *int32     `json:"dracoNormalQuantization,omitempty"`
	DracoGenericQuantization  *int32     `json:"dracoGenericQuantization,omitempty"`
}

type Model3DResult struct {
	JobID       uuid.UUID  `json:"jobID"`
	PartID      *uuid.UUID `json:"partID,omitempty"`
	ProductID   *uuid.UUID `json:"productID,omitempty"`
	Status      string     `json:"status"`
	DownloadURL *string    `json:"downloadURL,omitempty"` // non-nil only when status is "ready"
}
