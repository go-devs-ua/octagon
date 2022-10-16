package hash

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", h)
}
