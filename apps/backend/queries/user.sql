-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "user" WHERE username = $1 LIMIT 1;

-- name: GetUserByUsernameOrEmail :one
SELECT * FROM "user" WHERE username = $1 OR email = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO "user" (username, email, password_hash)
VALUES ($1, $2, $3)
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

-- name: UpdateUserWithPassword :exec
UPDATE "user"
SET
    username = $1,
    email = $2,
    password_hash = $3,
    email_verified = $4,
    banned = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

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
