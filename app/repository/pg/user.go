// Package repo represents adapter layer
// which enables interaction through a specific port
// and with a certain technology.
//
// in this case repo will act
// for CRUD operations with postgres.
package pg

import (
	"database/sql"

	"github.com/go-devs-ua/octagon/app/ent"
	"github.com/go-devs-ua/octagon/cfg"
)

type Repo struct{ *sql.DB }

func NewRepo(opt *cfg.Options) *Repo {
	return &Repo{}
}

// Add meth implements usecase.Repository
// interface without even knowing it
// that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r *Repo) Add(ent.User) error {
	// INSERT INTO "users" ...
	return nil
}
