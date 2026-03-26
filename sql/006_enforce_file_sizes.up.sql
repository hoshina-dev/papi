ALTER TABLE optimization_job_logs
    ALTER COLUMN source_file_size    SET NOT NULL,
    ALTER COLUMN processed_file_size SET NOT NULL,
    ALTER COLUMN compression_ratio   SET NOT NULL,
    ADD CONSTRAINT fk_optimization_job_logs_job_id
    FOREIGN KEY (job_id) REFERENCES part_3d_models(id);
