-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    ticket_id,
    quantity,
    total_price,
    additional_fees_id,
    payment_key,
    created_at,
    updated_at
)  
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetOrderByID :one
SELECT * FROM orders
WHERE id = $1
LIMIT 1;

-- name: GetOrderByPaymentKey :one
SELECT * FROM orders
WHERE payment_key = $1
LIMIT 1;