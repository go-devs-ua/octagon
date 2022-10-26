package entities

import (
	"github.com/google/uuid"
)

// ValidateUUID checks whether the input string is of UUID format.
func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}

	return nil
}
