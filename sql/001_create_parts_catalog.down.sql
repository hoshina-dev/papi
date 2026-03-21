-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS parts_inventory CASCADE;
DROP TABLE IF EXISTS part_categories CASCADE;
DROP TABLE IF EXISTS parts CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS manufacturers CASCADE;

-- Drop extensions if no longer needed
-- Note: Only drop if you're certain no other tables use them
-- DROP EXTENSION IF EXISTS "pg_trgm";
-- DROP EXTENSION IF EXISTS "uuid-ossp";
