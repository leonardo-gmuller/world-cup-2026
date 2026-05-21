-- name: GetStageWeight :one
SELECT * FROM stage_weights
WHERE stage = $1;

-- name: ListStageWeights :many
SELECT * FROM stage_weights
ORDER BY order_index ASC;