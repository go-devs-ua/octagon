package hash

import (
	"crypto/sha256"
	"fmt"
)

func Hash(someString string) string {
	h := sha256.Sum256([]byte(someString))
	return fmt.Sprintf("%x", h)
}
