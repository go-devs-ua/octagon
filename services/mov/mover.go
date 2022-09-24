// Package handler is responsible for logistics

package mov

import (
	"github.com/go-devs-ua/octagon/core"
)

// Mover takes care of the http requests,
// responses and validations.
type Mover struct {
	Agent core.Agency
}

// NewMover will return *Mover
// accepting Agency interface
func NewMover(ua core.Agency) *Mover {
	return &Mover{ua}
}
