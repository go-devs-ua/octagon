package rest

import (
	"regexp"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globs"
)

var idRegex = regexp.MustCompile(entities.IDMask)

func validateUUID(id string) error {
	if valid := idRegex.MatchString(id); !valid {
		return globs.ErrInvalidID
	}

	return nil
}
