-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS object_inventory_parts CASCADE;
DROP TABLE IF EXISTS object_inventory CASCADE;
