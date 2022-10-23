// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"fmt"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/pkg/hash"
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
func (r Repo) Add(user entities.User) (string, error) {
	var id string
	const sqlStatement = `INSERT INTO "user" (first_name, last_name, email, password)
						  VALUES ($1, $2, $3, $4) RETURNING id`
	if err := r.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, hash.SHA256(user.Password)).Scan(&id); err != nil {
		return "", fmt.Errorf("error inserting into database: %w", err)
	}

	return id, nil
}

// GetUserByID meth implements usecase.UserRepository logic
// finding user in DB by ID.
func (r Repo) Find(id string) (entities.PublicUser, error) {
	row := r.DB.QueryRow("SELECT id, first_name, last_name, email, created_at FROM user WHERE id=$1", id)
	var user entities.PublicUser

	if row.Err() == sql.ErrNoRows {
		return entities.PublicUser{}, fmt.Errorf("no user found in DB with such ID: %w", row.Err())
	}

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		return entities.PublicUser{}, fmt.Errorf("error while scanning row %w", err)
	}

	return user, nil
}
