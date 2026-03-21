-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- ============================================================================
-- Manufacturers Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS manufacturers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,
    country_of_origin CHAR(3), -- ISO 3166-1 alpha-3 country codes

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT check_iso_format
        CHECK (country_of_origin IS NULL OR country_of_origin ~ '^[A-Z]{3}$'),
    CONSTRAINT manufacturers_name_country_unique
        UNIQUE (name, country_of_origin)
);

CREATE INDEX IF NOT EXISTS idx_manufacturers_name ON manufacturers USING GIN (name gin_trgm_ops)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE manufacturers IS 'Catalog of part manufacturers';
COMMENT ON COLUMN manufacturers.country_of_origin IS 'ISO 3166-1 alpha-3 country code (e.g., USA, JPN, DEU)';

-- ============================================================================
-- Categories Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL UNIQUE,
    description TEXT,
    parent_id UUID REFERENCES categories(id) ON DELETE SET NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_categories_parent_id ON categories(parent_id)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE categories IS 'Hierarchical categories for parts organization';
COMMENT ON COLUMN categories.parent_id IS 'Reference to parent category for hierarchical structure';

-- ============================================================================
-- Parts Catalog Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Basic Information
    name TEXT NOT NULL,
    part_number TEXT NOT NULL,
    manufacturer_id UUID NOT NULL REFERENCES manufacturers(id) ON DELETE RESTRICT,
    description TEXT,

    -- Specifications
    temperature_stage TEXT, -- e.g., "ambient", "cryogenic", etc.
    specifications JSONB, -- Flexible storage for various specs (dimensions, weight, etc.)

    -- Media
    images TEXT[] NOT NULL DEFAULT '{}',

    -- Metadata
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT unique_manufacturer_part
        UNIQUE (manufacturer_id, part_number)
);

-- Indexes for parts table
CREATE INDEX IF NOT EXISTS idx_parts_manufacturer_id ON parts(manufacturer_id)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_name ON parts USING GIN (name gin_trgm_ops)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_part_number ON parts USING GIN (part_number gin_trgm_ops)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_created_at ON parts(created_at DESC)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_specifications ON parts USING GIN (specifications)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE parts IS 'Catalog of part types - this is the master list of all parts we work with';
COMMENT ON COLUMN parts.specifications IS 'JSON object for flexible specifications (dimensions, weight, material, etc.)';
COMMENT ON COLUMN parts.temperature_stage IS 'Operating temperature stage (ambient, cryogenic, etc.)';

-- ============================================================================
-- Part-Categories Junction Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS part_categories (
    part_id UUID NOT NULL REFERENCES parts(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (part_id, category_id)
);

CREATE INDEX IF NOT EXISTS idx_part_categories_part_id ON part_categories(part_id);
CREATE INDEX IF NOT EXISTS idx_part_categories_category_id ON part_categories(category_id);

COMMENT ON TABLE part_categories IS 'Many-to-many relationship between parts and categories';

-- ============================================================================
-- Parts Inventory Table
-- ============================================================================
CREATE TABLE IF NOT EXISTS parts_inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    part_id UUID NOT NULL REFERENCES parts(id) ON DELETE RESTRICT,
    serial_number TEXT NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT true,
    notes TEXT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT unique_part_serial UNIQUE (part_id, serial_number)
);

-- Indexes for parts_inventory table
CREATE INDEX IF NOT EXISTS idx_parts_inventory_part_id ON parts_inventory(part_id)
    WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_inventory_is_available ON parts_inventory(is_available)
    WHERE deleted_at IS NULL;

COMMENT ON TABLE parts_inventory IS 'Individual physical parts in inventory - each row represents a specific item';
COMMENT ON COLUMN parts_inventory.part_id IS 'Reference to the part type in the catalog';
COMMENT ON COLUMN parts_inventory.serial_number IS 'Serial number for this physical item (unique per part type)';
COMMENT ON COLUMN parts_inventory.is_available IS 'Whether this part is available for use (true = available, false = in use)';
