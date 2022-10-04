// Package main is key entry point
// of our awesome app
package main

import (
	"log"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	"github.com/go-devs-ua/octagon/app/transport/rest"
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
	opt := cfg.NewOptions()
	repo := pg.NewRepo(opt)

	userUsecase := usecase.NewUser(repo)
	userHandler := rest.NewUserHandler(userUsecase)

	server := rest.NewServer(userHandler)
	if err := server.Run(); err != nil {
		return err
	}

	return nil
}
