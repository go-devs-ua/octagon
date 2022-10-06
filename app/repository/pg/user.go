// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"fmt"

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
	const sqlStatement = `INSERT INTO "user" (first_name, last_name, email, password)
						  VALUES ($1, $2, $3, $4)`

	if _, err := r.DB.Exec(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return fmt.Errorf("error inserting into database: %w", err)
	}

	return nil
}
