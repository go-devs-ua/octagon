// Package main is key entry point
// of our awesome app.
package main

import (
	"fmt"
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	"github.com/go-devs-ua/octagon/app/transport/rest"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Cann't run server: %v", err)
	}
}

// Run will bind our layers all together.
func Run() error {
	config, err := cfg.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	logger, err := lgr.New(config.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	defer logger.Flush()

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		logger.Errorf("%+v", err)

		return fmt.Errorf("error connecting to database on host: %s, port: %s, with error: %w", config.DB.Host, config.DB.Port, err)
	}

	repo := pg.NewRepo(db)

	logger.Infof("Connection to database successfully created")

	handlers := rest.Handlers{
		UserHandler: rest.NewUserHandler(usecase.NewUser(repo), logger),
	}

	srv := rest.NewServer(config, handlers, logger)
	logger.Infof("Server starts on port:%s", config.Server.Port)

	if err := srv.Run(); err != nil {
		logger.Errorf("Failed loading server: %+v", err)

		return fmt.Errorf("error loading server: %w", err)
	}

	return nil
}
