ALTER TABLE optimization_job_logs
    ALTER COLUMN source_file_size    DROP NOT NULL,
    ALTER COLUMN processed_file_size DROP NOT NULL,
    ALTER COLUMN compression_ratio   DROP NOT NULL,
    DROP CONSTRAINT fk_optimization_job_logs_job_id;
