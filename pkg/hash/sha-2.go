package hash

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 function returns the checksum of the string using SHA256 hash algorithms.
func SHA256(someString string) string {
	h := sha256.Sum256([]byte(someString))

	return fmt.Sprintf("%x", h)
}
