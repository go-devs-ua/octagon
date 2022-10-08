// Package main is key entry point
// of our awesome app
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	rest "github.com/go-devs-ua/octagon/app/transport/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/go-devs-ua/octagon/migration"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run will bind our layers all together
func Run() error {
	logger := lgr.New()
	defer logger.Flush()

	opt, err := cfg.GetConfig()
	if err != nil {
		logger.Errorf("Failed to get config from .env: %+v\n", err)
		return err
	}

	db, err := connectDB(opt)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return err
	}

	repo := pg.NewRepo(db)
	logger.Infof("%s\n", "Connection to database successfully created")

	if err := migration.Migrate(db, logger); err != nil {
		logger.Errorf("Failed making migrations: %+v\n", err)
		return err
	}

	handlers := rest.Handlers{
		UserHandler: rest.NewUserHandler(usecase.NewUser(repo), logger),
	}

	srv := rest.NewServer(opt, handlers)
	if err := srv.Run(); err != nil {
		logger.Errorf("Failed loading server: %+v\n", err)
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
