-- name: CreateGroup :one
INSERT INTO groups (
    uuid, name, description, owner_id, invite_code
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetGroupByUUID :one
SELECT * FROM groups
WHERE uuid = $1
AND deleted_at IS NULL;

-- name: GetGroupByInviteCode :one
SELECT * FROM groups
WHERE invite_code = $1
AND deleted_at IS NULL;

-- name: ListGroupsByUserID :many
SELECT g.*
FROM groups g
INNER JOIN group_members gm ON gm.group_id = g.id
WHERE gm.user_id = $1
AND g.deleted_at IS NULL
AND gm.deleted_at IS NULL
ORDER BY g.created_at DESC;

-- name: CreateGroupMember :one
INSERT INTO group_members (
    uuid, group_id, user_id, role
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetGroupMember :one
SELECT *
FROM group_members
WHERE group_id = $1
AND user_id = $2
AND deleted_at IS NULL;

-- name: ListGroupMembers :many
SELECT *
FROM group_members
WHERE group_id = $1
AND deleted_at IS NULL
ORDER BY joined_at ASC;

-- name: DeleteGroupMember :exec
UPDATE group_members
SET deleted_at = NOW(),
    updated_at = NOW()
WHERE group_id = $1
AND user_id = $2
AND deleted_at IS NULL;