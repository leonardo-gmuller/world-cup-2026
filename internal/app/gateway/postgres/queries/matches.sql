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
SELECT
    m.*,
    ht.flag_url AS home_team_flag_url,
    at.flag_url AS away_team_flag_url
FROM matches m
LEFT JOIN teams ht ON ht.id = m.home_team_id
LEFT JOIN teams at ON at.id = m.away_team_id
WHERE m.deleted_at IS NULL
ORDER BY m.starts_at ASC;

-- name: ListMatchesByStage :many
SELECT * FROM matches
WHERE stage = $1
AND deleted_at IS NULL
ORDER BY starts_at ASC;

-- name: GetMatchByID :one
SELECT m.*, 
ht.flag_url AS home_team_flag_url, 
at.flag_url AS away_team_flag_url 
FROM 
matches m
LEFT JOIN teams ht ON ht.id = m.home_team_id
LEFT JOIN teams at ON at.id = m.away_team_id
WHERE m.id = $1
AND m.deleted_at IS NULL;

-- name: GetMatchByUUID :one
SELECT m.*, 
ht.flag_url AS home_team_flag_url, 
at.flag_url AS away_team_flag_url 
FROM 
matches m
LEFT JOIN teams ht ON ht.id = m.home_team_id
LEFT JOIN teams at ON at.id = m.away_team_id
WHERE m.uuid = $1
AND m.deleted_at IS NULL;

-- name: CountMatches :one
SELECT COUNT(*)
FROM matches m
WHERE m.deleted_at IS NULL;

-- name: GetNextMatch :one
SELECT
    m.*,
    ht.flag_url AS home_team_flag_url,
    at.flag_url AS away_team_flag_url
FROM matches m
LEFT JOIN teams ht ON ht.id = m.home_team_id
LEFT JOIN teams at ON at.id = m.away_team_id
WHERE m.deleted_at IS NULL
AND m.status = 'scheduled'
AND m.starts_at >= NOW()

ORDER BY m.starts_at ASC
LIMIT 1;

-- name: HasLiveMatches :one
SELECT EXISTS (
    SELECT 1
    FROM matches
    WHERE deleted_at IS NULL
    AND status = 'live'
);

-- name: FindMatchForLiveSync :one
SELECT *
FROM matches
WHERE LOWER(home_team_name) = LOWER(sqlc.arg(home_team_name)::text)
  AND LOWER(away_team_name) = LOWER(sqlc.arg(away_team_name)::text)
  AND starts_at BETWEEN
      sqlc.arg(starts_at)::timestamptz - INTERVAL '3 hours'
      AND
      sqlc.arg(starts_at)::timestamptz + INTERVAL '3 hours'
LIMIT 1;

-- name: UpdateLiveResult :one
UPDATE matches
SET
    api_football_id = sqlc.arg(api_football_id),
    home_score = sqlc.arg(home_score),
    away_score = sqlc.arg(away_score),
    status = sqlc.arg(status),
    result_source = 'api_football',
    last_live_sync_at = NOW(),
    updated_at = NOW()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetMatchByExternalID :one
SELECT *
FROM matches
WHERE external_id = $1
LIMIT 1;

-- name: ListMatchesToSyncLiveResults :many
SELECT *
FROM matches
WHERE deleted_at IS NULL
  AND (
    status = 'live'
    OR (
      status = 'scheduled'
      AND starts_at BETWEEN NOW() - INTERVAL '15 minutes'
                        AND NOW() + INTERVAL '15 minutes'
    )
  )
ORDER BY starts_at;