// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"

	"github.com/go-devs-ua/octagon/app/entities"
	_ "github.com/lib/pq"
)

// Repo wraps a database handle
type Repo struct {
	DB *sql.DB
}

// NewRepo will initialise new instance of Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

// Add meth implements usecase.UserRepository interface
// without even knowing it that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r *Repo) Add(user entities.User) error {
	// INSERT INTO "users" ...
	return nil
}
