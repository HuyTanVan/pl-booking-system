-- name: CreateIdempotency :one
INSERT INTO idempotency (
    user_id,
    idempotency_key,
    response_status_code,
    response_headers,
    response_body,
    created_at
)  
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetIdempotencyByUserID :one
SELECT * FROM idempotency
WHERE user_id = $1
LIMIT 1;

-- name: GetIdempotencyByIdempotencyKey :one
SELECT * FROM idempotency
WHERE idempotency_key = $1
LIMIT 1;
