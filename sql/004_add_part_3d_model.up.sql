CREATE TABLE IF NOT EXISTS part_3d_models (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    raw_url TEXT NOT NULL,

    status TEXT NOT NULL DEFAULT 'processing',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_part_3d_models_status ON part_3d_models(status);