// Package main is key entry point
// of our awesome app
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	rest "github.com/go-devs-ua/octagon/app/transport/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/migration"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Run will bind our layers all together
func Run() error {
	// TODO: Handle errors, migration, configs ...
	opt := cfg.NewOptions()

	db, err := connectDB(opt)
	if err != nil {
		return err
	}

	repo := pg.NewRepo(db)

	log.Println("Connected to DB")

	migration.Migrate(db)

	logic := usecase.NewUser(repo)
	srv := rest.NewServer(opt, logic)
	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}

func connectDB(opt cfg.Options) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		opt.DB.Host, opt.DB.Port, opt.DB.Username, opt.DB.Password, opt.DB.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
