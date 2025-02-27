-- name: CreateStadium :one
INSERT INTO stadiums (
    name,
    location,
    capacity,
    is_available
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListStadiums :many
SELECT * FROM stadiums
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetStadiumByID :one
SELECT * FROM stadiums
WHERE id = $1
LIMIT 1;

-- name: GetStadiumByName :one
SELECT * FROM stadiums
WHERE name = $1
LIMIT 1;

-- name: UpdateStadium :exec
UPDATE stadiums
SET name = coalesce(sqlc.narg('name'), name),
    location = coalesce(sqlc.narg('location'), location),
    capacity = coalesce(sqlc.narg('capacity'), capacity),
    is_available = coalesce(sqlc.narg('is_available'), is_available)
WHERE id=$1;