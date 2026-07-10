-- name: GetUserPasswordResetByTokenHash :one
SELECT * FROM user_password_reset WHERE token_hash = $1 LIMIT 1;

-- name: InsertUserPasswordReset :one
INSERT INTO user_password_reset (user_id, token_hash, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserPasswordResetsByUserID :exec
DELETE FROM user_password_reset WHERE user_id = $1;
