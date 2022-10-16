package hash

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", h)
}

/*
func hash(txt *string) error {
	h := sha256.Sum256([]byte(*txt))
	*txt = fmt.Sprintf("%x", h)
	return nil
}*/
