-- name: GetGroupRanking :many
SELECT
    u.id AS user_id,
    u.uuid AS user_uuid,
    u.name,
    COALESCE(SUM(p.points), 0)::float8 AS total_points,
    COUNT(p.id) AS predictions_count
FROM group_members gm
INNER JOIN users u ON u.id = gm.user_id
LEFT JOIN predictions p 
    ON p.user_id = u.id 
    AND p.group_id = gm.group_id
    AND p.deleted_at IS NULL
LEFT JOIN matches m
    ON m.id = p.match_id
WHERE gm.group_id = $1
AND gm.deleted_at IS NULL
AND u.deleted_at IS NULL
AND (sqlc.narg(stage)::text IS NULL OR m.stage = sqlc.narg(stage))
GROUP BY u.id, u.uuid, u.name
ORDER BY total_points DESC, u.name ASC;

-- name: ListGroupIDsByUserID :many
SELECT gm.group_id
FROM group_members gm
WHERE gm.user_id = $1
AND gm.deleted_at IS NULL;
