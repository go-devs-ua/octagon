// Package main is key entry point
// of our awesome app
package main

import (
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	rest "github.com/go-devs-ua/octagon/app/transport/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run will bind our layers all together
func Run() error {
	// TODO: Handle errors, migration, configs ...
	opt, err := cfg.GetConfig()
	if err != nil {
		return err
	}
	repo := pg.NewRepo(opt)
	logic := usecase.NewUser(repo)
	mux := rest.NewRouter()
	srv := rest.NewServer(logic, mux)
	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}
