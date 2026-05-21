-- name: UpsertMatch :one
INSERT INTO matches (
    uuid,
    external_id,
    stage,
    group_name,
    home_team_id,
    away_team_id,
    home_team_name,
    away_team_name,
    starts_at,
    home_score,
    away_score,
    status,
    winner_team_id,
    imported_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW()
)
ON CONFLICT (external_id)
DO UPDATE SET
    stage = EXCLUDED.stage,
    group_name = EXCLUDED.group_name,
    home_team_id = EXCLUDED.home_team_id,
    away_team_id = EXCLUDED.away_team_id,
    home_team_name = EXCLUDED.home_team_name,
    away_team_name = EXCLUDED.away_team_name,
    starts_at = EXCLUDED.starts_at,
    home_score = EXCLUDED.home_score,
    away_score = EXCLUDED.away_score,
    status = EXCLUDED.status,
    winner_team_id = EXCLUDED.winner_team_id,
    imported_at = NOW(),
    updated_at = NOW()
RETURNING *;

-- name: ListMatches :many
SELECT * FROM matches
WHERE deleted_at IS NULL
ORDER BY starts_at ASC;

-- name: ListMatchesByStage :many
SELECT * FROM matches
WHERE stage = $1
AND deleted_at IS NULL
ORDER BY starts_at ASC;

-- name: GetMatchByID :one
SELECT * FROM matches
WHERE id = $1
AND deleted_at IS NULL;

-- name: GetMatchByUUID :one
SELECT * FROM matches
WHERE uuid = $1
AND deleted_at IS NULL;
