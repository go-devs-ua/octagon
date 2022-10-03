// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"log"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/cfg"
)

// Repo wraps a database handle
type Repo struct{ sql.DB }

// NewRepo will initialise new instance of Repo
func NewRepo(opt cfg.Options) *Repo {
	return &Repo{}
}

// Add meth implements usecase.UserRepository interface
// without even knowing it that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r *Repo) Add(usr entities.User) error {
	sqlStatement := `
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id`
	err := r.QueryRow(sqlStatement, usr.FirstName, usr.LastName, usr.Email, usr.Password)
	if err != nil {
		log.Println(err)
	}

	return nil
}
