-- ============================================================================
-- Product Inventory Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS product_inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    product_id UUID NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    serial_number TEXT NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT true,
    notes TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT unique_product_serial UNIQUE (product_id, serial_number)
);

CREATE INDEX IF NOT EXISTS idx_product_inventory_product_id ON product_inventory(product_id)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_product_inventory_is_available ON product_inventory(is_available)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE product_inventory IS 'Individual built instances of products (e.g., specific built "Weather Station #1")';
COMMENT ON COLUMN product_inventory.product_id IS 'Reference to the product type (what was built)';
COMMENT ON COLUMN product_inventory.serial_number IS 'Serial number for this specific built product (unique per product type)';
COMMENT ON COLUMN product_inventory.is_available IS 'Whether this built product is available (true = available, false = in use)';

-- ============================================================================
-- Product Inventory Parts (Tracks which physical parts are in this product)
-- ============================================================================
CREATE TABLE IF NOT EXISTS product_inventory_parts (
    product_inventory_id UUID NOT NULL REFERENCES product_inventory(id) ON DELETE CASCADE,
    part_inventory_id UUID NOT NULL REFERENCES parts_inventory(id) ON DELETE RESTRICT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (product_inventory_id, part_inventory_id)
);

CREATE INDEX IF NOT EXISTS idx_product_inventory_parts_product_inventory_id
    ON product_inventory_parts(product_inventory_id);
CREATE INDEX IF NOT EXISTS idx_product_inventory_parts_part_inventory_id
    ON product_inventory_parts(part_inventory_id);

COMMENT ON TABLE product_inventory_parts IS 'Optional tracking of which specific physical parts went into building this product instance. Constraints are loose - you can create product_inventory without specifying any parts.';
COMMENT ON COLUMN product_inventory_parts.product_inventory_id IS 'The built product instance';
COMMENT ON COLUMN product_inventory_parts.part_inventory_id IS 'A specific physical part used in this product';
