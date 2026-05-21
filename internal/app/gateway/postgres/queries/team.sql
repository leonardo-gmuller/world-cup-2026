-- name: UpsertTeam :one
INSERT INTO teams (
    uuid,
    external_id,
    name,
    short_name,
    code,
    flag_url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
ON CONFLICT (name)
DO UPDATE SET
    external_id = EXCLUDED.external_id,
    short_name = EXCLUDED.short_name,
    code = EXCLUDED.code,
    flag_url = EXCLUDED.flag_url,
    updated_at = NOW()
RETURNING *;