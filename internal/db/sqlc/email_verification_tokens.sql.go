// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: email_verification_tokens.sql

package db

import (
	"context"
)

const createEmailVerificationToken = `-- name: CreateEmailVerificationToken :one
INSERT INTO email_verification_tokens (
    user_id,
    token
)  
VALUES ($1, $2)
RETURNING id, user_id, token
`

type CreateEmailVerificationTokenParams struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

func (q *Queries) CreateEmailVerificationToken(ctx context.Context, arg CreateEmailVerificationTokenParams) (EmailVerificationToken, error) {
	row := q.db.QueryRowContext(ctx, createEmailVerificationToken, arg.UserID, arg.Token)
	var i EmailVerificationToken
	err := row.Scan(&i.ID, &i.UserID, &i.Token)
	return i, err
}

const getEVTokenByUserID = `-- name: GetEVTokenByUserID :one
SELECT id, user_id, token FROM email_verification_tokens
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetEVTokenByUserID(ctx context.Context, userID int32) (EmailVerificationToken, error) {
	row := q.db.QueryRowContext(ctx, getEVTokenByUserID, userID)
	var i EmailVerificationToken
	err := row.Scan(&i.ID, &i.UserID, &i.Token)
	return i, err
}
