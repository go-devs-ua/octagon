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
	opt := cfg.NewOptions()
	repo := pg.NewRepo(opt)

	userLogic := usecase.NewUser(repo)
	user := rest.NewUserHandler(userLogic)

	mux := rest.NewRouter(user.Map()...) // here we can put other entities like admin.Map()...

	srv := rest.NewServer(mux)
	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}
