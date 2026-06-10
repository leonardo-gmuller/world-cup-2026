-- name: UpsertPrediction :one
INSERT INTO predictions (
    uuid, group_id, user_id, match_id, home_score, away_score
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (group_id, user_id, match_id)
DO UPDATE SET
    home_score = EXCLUDED.home_score,
    away_score = EXCLUDED.away_score,
    calculated = FALSE,
    calculated_at = NULL,
    updated_at = NOW()
WHERE predictions.calculated = FALSE
AND predictions.deleted_at IS NULL
RETURNING *;

-- name: GetPrediction :one
SELECT * FROM predictions
WHERE group_id = $1
AND user_id = $2
AND match_id = $3
AND deleted_at IS NULL;

-- name: ListPredictionsByUserAndGroup :many
SELECT p.*,
        m.uuid AS match_uuid,
        g.uuid AS group_uuid
FROM predictions p
JOIN matches m ON p.match_id = m.id
JOIN groups g ON p.group_id = g.id
WHERE p.group_id = $1
AND p.user_id = $2
AND p.deleted_at IS NULL
ORDER BY p.created_at DESC;

-- name: ListPredictionsByMatch :many
SELECT * FROM predictions
WHERE match_id = $1
AND deleted_at IS NULL;

-- name: GetPredictionByID :one
SELECT *
FROM predictions
WHERE uuid = $1
AND deleted_at IS NULL;

-- name: UpdatePredictionPoints :exec
UPDATE predictions
SET points = $2,
    calculated = $3,
    calculated_at = NOW(),
    updated_at = NOW()
WHERE id = $1
AND deleted_at IS NULL;

-- name: ListFinishedOrLiveMatchesToCalculate :many
SELECT *
FROM matches m
WHERE m.status IN ('finished', 'live')
AND m.deleted_at IS NULL
AND EXISTS (
    SELECT 1
    FROM predictions p
    WHERE p.match_id = m.id
    AND p.calculated = FALSE
    AND p.deleted_at IS NULL
)
ORDER BY m.starts_at ASC;

-- name: CountPredictionsByUserID :one
SELECT COUNT(*)
FROM predictions p
WHERE p.user_id = $1
AND p.deleted_at IS NULL;

-- name: ListPredictionRemindersByUserID :many
SELECT
    m.id AS match_id,
    gm.group_id,
    g.name AS group_name,
    m.home_team_name,
    m.away_team_name,
    ht.flag_url AS home_team_flag_url,
    at.flag_url AS away_team_flag_url,
    m.starts_at
FROM group_members gm
INNER JOIN groups g
    ON g.id = gm.group_id
    AND g.deleted_at IS NULL
INNER JOIN matches m
    ON m.deleted_at IS NULL
    AND m.status = 'scheduled'
    AND m.starts_at > NOW() + INTERVAL '5 minutes'
    AND m.starts_at <= NOW() + INTERVAL '48 hours'
LEFT JOIN teams ht ON ht.id = m.home_team_id
LEFT JOIN teams at ON at.id = m.away_team_id
LEFT JOIN predictions p
    ON p.group_id = gm.group_id
    AND p.user_id = gm.user_id
    AND p.match_id = m.id
    AND p.deleted_at IS NULL
WHERE gm.user_id = $1
AND gm.deleted_at IS NULL
AND p.id IS NULL
ORDER BY m.starts_at ASC, g.name ASC;