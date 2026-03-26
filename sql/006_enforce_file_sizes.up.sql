-- Make file size and compression ratio non-nullable in optimization_job_logs.
-- source_file_size and processed_file_size are now required fields in the
-- worker webhook contract, so NULL values should never occur going forward.

ALTER TABLE optimization_job_logs
    ALTER COLUMN source_file_size    SET NOT NULL,
    ALTER COLUMN processed_file_size SET NOT NULL,
    ALTER COLUMN compression_ratio   SET NOT NULL;
