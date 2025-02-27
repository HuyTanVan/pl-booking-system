-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name, 
    email, 
    password, 
    phone_number,
    is_active,
    created_at,
    updated_at
)  
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
SET first_name = coalesce(sqlc.narg('first_name'), first_name),
    last_name = coalesce(sqlc.narg('last_name'), last_name),
    password = coalesce(sqlc.narg('password'), password),
    phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
    is_active = coalesce(sqlc.narg('is_active'), is_active)
WHERE id=$1;