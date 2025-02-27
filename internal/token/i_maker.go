package token

import "time"

type IMaker interface {
	CreateToken(email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
