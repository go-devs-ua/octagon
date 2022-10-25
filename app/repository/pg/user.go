// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/pkg/hash"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // Standart blanc import for pq.
)

var (
	ErrDuplicateEmail = errors.New("email is already taken")
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
		var pqErr = new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == uniqueViolationErrCode {
			return "", ErrDuplicateEmail
		}

		return "", fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}

func (r Repo) IsUserExists(user entities.UserID) (bool, error) {
	var err error
	var id string

	const sqlStatement = `SELECT id FROM "user" WHERE id = $1 AND deleted_at IS NULL`

	if err = r.DB.QueryRow(sqlStatement, user.ID).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("internal error while querying: %w", err)
	}

	return true, nil
}

func (r Repo) Delete(user entities.UserID) error {
	const sqlStatement = `UPDATE "user" SET deleted_at = NOW() WHERE id = $1`

	if _, err := r.DB.Exec(sqlStatement, user.ID); err != nil {
		return fmt.Errorf("error inserting into database: %w", err)
	}

	return nil
}
