DELETE FROM optimization_job_logs;
DELETE FROM part_3d_models;

ALTER TABLE part_3d_models RENAME TO model_3d;

ALTER TABLE model_3d
    ADD COLUMN part_id UUID REFERENCES parts(id),
    ADD COLUMN product_id UUID REFERENCES products(id),
    ADD CONSTRAINT check_model_3d_owner CHECK (
        part_id IS NOT NULL OR product_id IS NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_model_3d_part_id ON model_3d(part_id);
CREATE INDEX IF NOT EXISTS idx_model_3d_product_id ON model_3d(product_id);