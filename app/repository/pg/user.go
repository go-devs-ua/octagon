// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/cfg"
	_ "github.com/lib/pq"
)

// Repo wraps a database handle
type Repo struct {
	DB *sql.DB
}

// NewRepo will initialise new instance of Repo
func NewRepo(opt cfg.Options) *Repo {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		opt.DBConfig.Host, opt.DBConfig.Port, opt.DBConfig.Username, opt.DBConfig.password, opt.DBConfig.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		log.Println(fmt.Errorf("sql.Open(): %w", err))
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Println(fmt.Errorf("db.Ping(): %w", err))
		return nil
	}

	log.Println("Connected to database")

	return &Repo{
		DB: db,
	}
}

// Add meth implements usecase.UserRepository interface
// without even knowing it that allow us to decouple our layers
// and will make our app flexible and maintainable.
func (r *Repo) Add(usr entities.User) error {
	// INSERT INTO "users" ...
	return nil
}
