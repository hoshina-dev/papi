-- Create read-write role for the application
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'papi_rw') THEN
        CREATE ROLE papi_rw;
    END IF;
END
$$;

-- Grant connection privileges
GRANT CONNECT ON DATABASE postgres TO papi_rw;

-- Grant schema usage
GRANT USAGE ON SCHEMA public TO papi_rw;

-- Grant table privileges (will apply to future tables via defaults)
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO papi_rw;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO papi_rw;

-- Grant sequence privileges
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO papi_rw;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT USAGE, SELECT ON SEQUENCES TO papi_rw;

-- Grant function execution
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO papi_rw;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT EXECUTE ON FUNCTIONS TO papi_rw;
