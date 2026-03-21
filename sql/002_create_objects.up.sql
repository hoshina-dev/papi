-- ============================================================================
-- Objects Table (Bill of Materials)
-- ============================================================================
CREATE TABLE IF NOT EXISTS objects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,
    description TEXT,
    version TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_objects_name ON objects USING GIN (name gin_trgm_ops)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE objects IS 'Objects that are built from multiple parts (e.g., assemblies, projects)';
COMMENT ON COLUMN objects.version IS 'Version of this object design (e.g., v1.0, v2.3)';

-- ============================================================================
-- Object Parts (BOM - Bill of Materials)
-- ============================================================================
CREATE TABLE IF NOT EXISTS object_parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    object_id UUID NOT NULL REFERENCES objects(id) ON DELETE CASCADE,
    part_id UUID NOT NULL REFERENCES parts(id) ON DELETE RESTRICT,
    quantity INT NOT NULL CHECK (quantity > 0),

    notes TEXT, -- Optional notes about this part in this object

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_object_part UNIQUE (object_id, part_id)
);

CREATE INDEX IF NOT EXISTS idx_object_parts_object_id ON object_parts(object_id);
CREATE INDEX IF NOT EXISTS idx_object_parts_part_id ON object_parts(part_id);

COMMENT ON TABLE object_parts IS 'Bill of Materials - which parts are needed to build an object and how many';
COMMENT ON COLUMN object_parts.quantity IS 'Number of this part needed to build the object';
