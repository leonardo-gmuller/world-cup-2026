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