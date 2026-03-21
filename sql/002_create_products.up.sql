-- ============================================================================
-- Products Table (Bill of Materials)
-- ============================================================================
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,
    version TEXT,
    description TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_products_name ON products USING GIN (name gin_trgm_ops)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE products IS 'Products that are built from multiple parts (e.g., assemblies, projects)';

-- ============================================================================
-- Product Parts (BOM - Bill of Materials)
-- ============================================================================
CREATE TABLE IF NOT EXISTS product_parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    part_id UUID NOT NULL REFERENCES parts(id) ON DELETE RESTRICT,
    quantity INT NOT NULL CHECK (quantity > 0),

    notes TEXT, -- Optional notes about this part in this product

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_product_part UNIQUE (product_id, part_id)
);

CREATE INDEX IF NOT EXISTS idx_product_parts_product_id ON product_parts(product_id);
CREATE INDEX IF NOT EXISTS idx_product_parts_part_id ON product_parts(part_id);

COMMENT ON TABLE product_parts IS 'Bill of Materials - which parts are needed to build a product and how many';
COMMENT ON COLUMN product_parts.quantity IS 'Number of this part needed to build the product';
