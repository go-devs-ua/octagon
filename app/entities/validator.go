package entities

import (
	"regexp"

	"github.com/go-devs-ua/octagon/app/globals"
)

const idMask = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$" //nolint:gosec,lll // "Potential hardcoded credentials" regexp can`t be changed.

var idRegex = regexp.MustCompile(idMask)

// ValidateUUID checks whether the input string is of UUID format.
func ValidateUUID(id string) error {
	if valid := idRegex.MatchString(id); !valid {
		return globals.ErrInvalidID
	}

	return nil
}
