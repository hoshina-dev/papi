-- Reverses 999_mock_data.up.sql: wipe all catalog / inventory / product data

TRUNCATE TABLE
    product_inventory_parts,
    product_inventory,
    product_parts,
    products,
    parts_inventory,
    part_categories,
    parts,
    categories,
    manufacturers,
    optimization_job_logs,
    model_3d
RESTART IDENTITY CASCADE;
