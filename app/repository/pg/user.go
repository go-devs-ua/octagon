// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

// Add meth implements usecase.UserRepository interface
// without even knowing it that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r *Repo) Add(usr entities.User) error {
	// INSERT INTO "users" ...
	return nil
}
