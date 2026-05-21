-- name: CreateJobControl :exec
INSERT INTO jobs_control (job)
VALUES ($1)
ON CONFLICT DO NOTHING;

-- name: UpdateJobControl :exec
UPDATE jobs_control
SET last_success_run = NOW(),
    updated_at = NOW()
WHERE job = $1;

-- name: GetJobControl :one
SELECT job, last_success_run, is_enabled
FROM jobs_control
WHERE job = $1;