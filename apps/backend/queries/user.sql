-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "user" WHERE username = $1 LIMIT 1;

-- name: GetUserByUsernameOrEmail :one
SELECT * FROM "user" WHERE username = $1 OR email = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO "user" (username, email, email_verified, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET
    username = $1,
    email = $2,
    email_verified = $3,
    banned = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $5;

-- name: UpdateUserPassword :exec
UPDATE "user"
SET
    password_hash = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: VerifyUserEmail :exec
UPDATE "user"
SET
    email_verified = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: BanUser :exec
UPDATE "user"
SET
    banned = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
