CREATE TABLE IF NOT EXISTS part_3d_models (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    raw_url      TEXT NOT NULL,
    processed_key TEXT,

    status TEXT NOT NULL DEFAULT 'processing',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON COLUMN part_3d_models.processed_key IS 'S3 key for the optimized/processed 3D model file';

CREATE INDEX IF NOT EXISTS idx_part_3d_models_status ON part_3d_models(status);