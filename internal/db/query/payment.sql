-- name: CreatePayment :one
INSERT INTO payments (
    payment_key,
    amount,
    payment_method_id,
    status,
    date,
    updated_at
)  
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPaymentByID :one
SELECT * FROM payments
WHERE id = $1
LIMIT 1;

-- name: GetPaymentByPaymentKey :one
SELECT * FROM payments
WHERE payment_key = $1
LIMIT 1;

-- name: UpdatePaymentStatusByPaymentKey :exec
UPDATE payments
SET status = $2
WHERE payment_key=$1;