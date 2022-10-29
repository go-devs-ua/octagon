// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/go-devs-ua/octagon/pkg/hash"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // Standart blanc import for pq.
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

// AddUser method implements storing the user in the database.
func (r Repo) AddUser(user entities.User) (string, error) {
	var id string

	const sqlStatement = `INSERT INTO "user" (first_name, last_name, email, password)
						  VALUES ($1, $2, $3, $4) RETURNING id`

	if err := r.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, hash.SHA256(user.Password)).Scan(&id); err != nil {
		var pqErr = new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", globals.ErrDuplicateEmail
		}

		return "", fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}

// FindUser method implements logic of finding user in the database by ID.
func (r Repo) FindUser(id string) (*entities.User, error) {
	var user entities.User

	const sqlStatement = `SELECT id, first_name, last_name, email, created_at FROM "user" WHERE id=$1`

	if err := r.DB.QueryRow(sqlStatement, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, globals.ErrNotFound
		}

		return nil, fmt.Errorf("internal error while scanning row: %w", err)
	}

	return &user, nil
}

func setDefaultIfValueIsZero(name string, value any, defaultValue any) sql.NamedArg {
	if reflect.ValueOf(value).IsZero() {
		value = defaultValue
	}

	return sql.Named(name, value)
}

// GetUsers retrieves list of users from database.
func (r Repo) GetUsers(offset, limit, sort string) ([]entities.User, error) {
	const sqlStatement = `
			SELECT id, first_name, last_name, created_at
			FROM "user" 
			WHERE deleted_at IS NULL
			ORDER BY $1
			OFFSET $2 
			LIMIT $3 
	`
	var null sql.NullString

	rows, err := r.DB.Query(sqlStatement,
		setDefaultIfValueIsZero("sort", sort, defaultSort),
		setDefaultIfValueIsZero("offset", offset, null),
		setDefaultIfValueIsZero("limit", limit, null),
	)

	if err != nil {
		var pqErr = new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeInvalidTextRepresentation {
			return nil, fmt.Errorf("error occured while inserting query arguments: %w", globals.ErrBadQuery)
		}

		return nil, fmt.Errorf("error occurred while executing query: %w", err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("error occurred while scaning object from query: %w", err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %w", err)
	}

	return users, nil
}
