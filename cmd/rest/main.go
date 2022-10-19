// Package main is key entry point
// of our awesome app
package main

import (
	"fmt"
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	rest "github.com/go-devs-ua/octagon/app/transport/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Cann't run server: %v", err)
	}
}

// Run will bind our layers all together
func Run() error {
	config, err := cfg.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config from .env: %+v", err)
	}

	logger, err := lgr.New(config.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	defer logger.Flush()

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		logger.Errorf("%+v", err)
		return err
	}

	repo := pg.NewRepo(db)
	logger.Infof("Connection to database successfully created")

	handlers := rest.Handlers{
		UserHandler: rest.NewUserHandler(usecase.NewUser(repo), logger),
	}

	srv := rest.NewServer(config, handlers, logger)
	if err := srv.Run(); err != nil {
		logger.Errorf("Failed loading server: %+v", err)
		return err
	}

	return nil
}
