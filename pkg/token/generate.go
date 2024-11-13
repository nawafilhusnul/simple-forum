package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken(userID int64) string {
	b := make([]byte, 255)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
