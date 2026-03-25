-- Drop optimization_job_logs table
DROP INDEX IF EXISTS idx_optimization_job_logs_duration;
DROP INDEX IF EXISTS idx_optimization_job_logs_completed_at;
DROP INDEX IF EXISTS idx_optimization_job_logs_created_at;
DROP INDEX IF EXISTS idx_optimization_job_logs_status;
DROP INDEX IF EXISTS idx_optimization_job_logs_job_id;
DROP TABLE IF EXISTS optimization_job_logs;

