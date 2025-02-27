package utils

import (
	"errors"
	"math/rand"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// HashPassword returns the bcrypt hash of the password
func RandomString(length int8) (string, error) {
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	if len(b) != 32 {
		return string(b), errors.New("cannot create a random string of length 32")
	}
	return string(b), nil
}
