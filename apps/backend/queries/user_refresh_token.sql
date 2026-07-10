-- name: GetUserRefreshTokenByRefreshTokenHash :one
SELECT * FROM user_refresh_token WHERE token_hash = $1;

-- name: InsertUserRefreshToken :one
INSERT INTO user_refresh_token (user_id, token_hash, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserRefreshTokenByRefreshTokenHash :exec
DELETE FROM user_refresh_token WHERE token_hash = $1;

-- name: DeleteUserRefreshTokensByUserID :exec
DELETE FROM user_refresh_token WHERE user_id = $1;

