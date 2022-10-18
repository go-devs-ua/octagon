// Package main is key entry point
// of our awesome app
package main

import (
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	rest "github.com/go-devs-ua/octagon/app/transport/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
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

	config, err := cfg.GetConfig()
	if err != nil {
		logger.Errorf("Failed to get config from .env: %+v\n", err)
		return err
	}

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return err
	}

	repo := pg.NewRepo(db)
	logger.Infof("Connection to database successfully created\n")

	handlers := rest.Handlers{
		UserHandler: rest.NewUserHandler(usecase.NewUser(repo), logger),
	}

	srv := rest.NewServer(config, handlers, logger)
	if err := srv.Run(); err != nil {
		logger.Errorf("Failed loading server: %+v\n", err)
		return err
	}

	return nil
}
