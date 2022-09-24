// Package keeper will act for CRUD operations with database.

package repo

import (
	"database/sql"

	"github.com/go-devs-ua/octagon/cfg"
)

type Repo struct{ *sql.DB }

// NewRepo will initialise database using *cfg.Options.
func NewRepo(opt *cfg.Options) *Repo {
	return &Repo{}
}
