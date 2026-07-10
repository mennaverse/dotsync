-- name: GetUserVerificationTokenByHash :one
SELECT * FROM user_verification_token WHERE token_hash = $1 LIMIT 1;

-- name: InsertUserVerificationToken :one
INSERT INTO user_verification_token (user_id, token_hash, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserVerificationTokensByUserID :exec
DELETE FROM user_verification_token WHERE user_id = $1;
