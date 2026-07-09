-- name: FindUserByID :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM "user" WHERE email = $1 LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE username = $1 LIMIT 1;

-- name: FindUserByUsernameOrEmail :one
SELECT * FROM "user" WHERE username = $1 OR email = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO "user" (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE "user"
SET
    username = $1,
    email = $2,
    email_verified = $3,
    banned = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $5
RETURNING *;

-- name: UpdateUserWithPassword :one
UPDATE "user"
SET
    username = $1,
    email = $2,
    password_hash = $3,
    email_verified = $4,
    banned = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $6
RETURNING *;

-- name: VerifyUserEmail :one
UPDATE "user"
SET
    email_verified = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: BanUser :one
UPDATE "user"
SET
    banned = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
