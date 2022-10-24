package rest

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidID = errors.New("invalid ID")
	idRegex      = regexp.MustCompile(idMask)
)

func validateID(id string) error {
	if valid := idRegex.MatchString(id); !valid {
		return ErrInvalidID
	}

	return nil
}
