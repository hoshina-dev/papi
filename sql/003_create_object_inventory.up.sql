-- ============================================================================
-- Object Inventory Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS object_inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    object_id UUID NOT NULL REFERENCES objects(id) ON DELETE RESTRICT,
    serial_number TEXT NOT NULL,
    owner TEXT,
    notes TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT unique_object_serial UNIQUE (object_id, serial_number)
);

CREATE INDEX IF NOT EXISTS idx_object_inventory_object_id ON object_inventory(object_id)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_object_inventory_owner ON object_inventory(owner)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE object_inventory IS 'Individual built instances of objects (e.g., specific built "Lego Death Star #1")';
COMMENT ON COLUMN object_inventory.object_id IS 'Reference to the object type (what was built)';
COMMENT ON COLUMN object_inventory.serial_number IS 'Serial number for this specific built object (unique per object type)';
COMMENT ON COLUMN object_inventory.owner IS 'Current owner or location of this built object';

-- ============================================================================
-- Object Inventory Parts (Tracks which physical parts are in this object)
-- ============================================================================
CREATE TABLE IF NOT EXISTS object_inventory_parts (
    object_inventory_id UUID NOT NULL REFERENCES object_inventory(id) ON DELETE CASCADE,
    part_inventory_id UUID NOT NULL REFERENCES parts_inventory(id) ON DELETE RESTRICT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (object_inventory_id, part_inventory_id)
);

CREATE INDEX IF NOT EXISTS idx_object_inventory_parts_object_inventory_id
    ON object_inventory_parts(object_inventory_id);
CREATE INDEX IF NOT EXISTS idx_object_inventory_parts_part_inventory_id
    ON object_inventory_parts(part_inventory_id);

COMMENT ON TABLE object_inventory_parts IS 'Optional tracking of which specific physical parts went into building this object instance. Constraints are loose - you can create object_inventory without specifying any parts.';
COMMENT ON COLUMN object_inventory_parts.object_inventory_id IS 'The built object instance';
COMMENT ON COLUMN object_inventory_parts.part_inventory_id IS 'A specific physical part used in this object';
