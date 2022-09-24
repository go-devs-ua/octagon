// Package keeper will act for CRUD operations with database.

package repo

import (
	"database/sql"

	"github.com/delveper/heroes/core"
)

type User = core.User

type Repo struct {
	db *sql.DB
}

// NewRepo will initialise database.
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

func (kpr *Repo) Add(User) error {
	return nil
}
