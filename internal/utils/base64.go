package utils

import (
	"encoding/base64"
)

// EncodeToBase64 encodes a string to Base64
func EncodeToBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
func EncodeByteToBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// DecodeFromBase64 decodes a Base64 string
func DecodeFromBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// func DecodeFromByte(encoded []byte) ([]byte, error) {
// 	return base64.StdEncoding.DecodeString(encoded)
// }
