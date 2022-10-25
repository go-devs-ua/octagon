package rest

import (
	"errors"
	"regexp"

	"github.com/go-devs-ua/octagon/app/entities"
)

var (
	ErrInvalidID = errors.New("invalid ID")
	idRegex      = regexp.MustCompile(entities.IdMask)
)

func validateID(id string) error {
	if valid := idRegex.MatchString(id); !valid {
		return ErrInvalidID
	}

	return nil
}
