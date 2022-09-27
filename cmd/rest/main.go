// Package main is key entry point
// of our awesome app
package main

import (
	"log"

	"github.com/go-devs-ua/octagon/app/repo/pg"
	rest "github.com/go-devs-ua/octagon/app/trans/http"
	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/go-devs-ua/octagon/cfg"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	// TODO: Handle errors gracefully
	opt := cfg.NewOptions()
	repo := pg.NewRepo(opt)
	logic := usecase.NewUseCase(repo)
	mux := rest.NewRouter()
	srv := rest.NewServer(logic, mux)
	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}
