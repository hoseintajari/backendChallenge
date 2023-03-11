package Utility

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	hashed := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hashed)
}
