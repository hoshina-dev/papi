DROP INDEX IF EXIST idx_model_3d_product_id;
DROP INDEX IF EXIST idx_model_3d_part_id;

ALTER TABLE model_3d
    DROP CONSTRAINT check_model_3d_owner,
    DROP COLUMN product_id,
    DROP COLUMN part_id;

ALTER TABLE model_3d RENAME TO part_3d_models;