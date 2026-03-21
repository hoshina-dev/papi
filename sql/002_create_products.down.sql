-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS object_parts CASCADE;
DROP TABLE IF EXISTS objects CASCADE;
