package util

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func HashTokenWithSha256(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func VerifyToken(token, hashedToken string) bool {
	hashedInput := HashTokenWithSha256(token)
	return subtle.ConstantTimeCompare([]byte(hashedInput), []byte(hashedToken)) == 1
}
