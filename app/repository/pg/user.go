// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"fmt"
	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/pkg/hash"
	_ "github.com/lib/pq" // Standart blanc import for pq.
	"strings"
)

// Repo wraps a database handle.
type Repo struct {
	DB *sql.DB
}

// NewRepo will initialise new instance of Repo.
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

// Add meth implements usecase.UserRepository interface
// without even knowing it that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r Repo) Add(user entities.User) (string, error) {
	var id string

	const sqlStatement = `INSERT INTO "user" (first_name, last_name, email, password)
						  VALUES ($1, $2, $3, $4) RETURNING id`

	if err := r.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, hash.SHA256(user.Password)).Scan(&id); err != nil {
		if strings.Contains(err.Error(), "unique_user_email") { //FIXME by Rusli4 for Andrii
			return "", usecase.ErrDuplicateEmail
		}

		return "", fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}
