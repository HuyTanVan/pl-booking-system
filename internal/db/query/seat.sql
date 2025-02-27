-- name: CreateSeat :one
INSERT INTO seats (
    stadium_id,
    block,
    row,
    seat_column,
    seat_available
)  
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSeatByID :one
SELECT * FROM seats
WHERE id = $1 LIMIT 1;

-- name: ListSeats :many
SELECT * FROM seats
WHERE seats.stadium_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListSeatsWithDetails :many
SELECT s.*, 
       stadium.name AS stadium_name
FROM seats s
LEFT JOIN stadiums AS stadium ON s.stadium_id = $1
LIMIT $2
OFFSET $3;

-- -- name: UpdateUser :exec
-- UPDATE payment_methods
-- SET first_name = coalesce(sqlc.narg('first_name'), first_name),
--     last_name = coalesce(sqlc.narg('last_name'), last_name),
--     password = coalesce(sqlc.narg('password'), password),
--     phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
--     is_active = coalesce(sqlc.narg('is_active'), is_active)
-- WHERE id=$1;

-- name: UpdateSeatAvailable :exec
UPDATE seats
SET seat_available = $1
WHERE id = $2;