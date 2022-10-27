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

func (r Repo) GetUsers(offset, limit, sort string) ([]entities.User, error) {
	const sqlStatement = `
			SELECT id, first_name, last_name, created_at
			FROM "user" 
			WHERE deleted_at IS NULL
			ORDER BY $1 
			OFFSET $2 
			LIMIT $3;
	`

	rows, err := r.DB.Query(sqlStatement, sort, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scaning object from query: %w", err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured during iteration: %w", err)
	}

	return users, nil
}
