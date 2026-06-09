-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (
    user_id,
    token,
    expires_at
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetPasswordResetTokenByToken :one
SELECT *
FROM password_reset_tokens
WHERE token = $1
LIMIT 1;

-- name: InvalidatePasswordResetTokensByUserID :exec
UPDATE password_reset_tokens
SET
    used = TRUE,
    used_at = NOW()
WHERE
    user_id = $1
    AND used = FALSE;

-- name: UsePasswordResetToken :exec
UPDATE password_reset_tokens
SET
    used = TRUE,
    used_at = NOW()
WHERE id = $1;

-- name: DeleteExpiredPasswordResetTokens :exec
DELETE FROM password_reset_tokens
WHERE expires_at < NOW();

-- name: ListActivePasswordResetTokensByUserID :many
SELECT *
FROM password_reset_tokens
WHERE
    user_id = $1
    AND used = FALSE
    AND expires_at > NOW()
ORDER BY created_at DESC;