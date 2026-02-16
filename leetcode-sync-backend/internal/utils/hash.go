package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashCode(code string) string {
	hash := sha256.Sum256([]byte(code))
	return hex.EncodeToString(hash[:])
}
