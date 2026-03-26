ALTER TABLE optimization_job_logs
    ALTER COLUMN source_file_size    DROP NOT NULL,
    ALTER COLUMN processed_file_size DROP NOT NULL,
    ALTER COLUMN compression_ratio   DROP NOT NULL;
