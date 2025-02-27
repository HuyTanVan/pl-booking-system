-- name: CreateEmailVerificationToken :one
INSERT INTO email_verification_tokens (
    user_id,
    token
)  
VALUES ($1, $2)
RETURNING *;

-- name: GetEVTokenByUserID :one
SELECT * FROM email_verification_tokens
WHERE user_id = $1 LIMIT 1;
