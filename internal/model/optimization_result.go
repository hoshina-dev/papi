package model

import "github.com/google/uuid"

type Optimize3DParams struct {
	PartID                    *uuid.UUID
	ProductID                 *uuid.UUID
	SourceURL                 string
	DracoCompressionLevel     *int32
	DracoPositionQuantization *int32
	DracoTexcoordQuantization *int32
	DracoNormalQuantization   *int32
	DracoGenericQuantization  *int32
}

type JobResult struct {
	JobID       uuid.UUID
	PartID      *uuid.UUID
	ProductID   *uuid.UUID
	Status      string
	DownloadURL *string // non-nil only when status is "ready"
}
