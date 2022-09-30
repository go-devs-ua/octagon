package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

// UserUsecase represents User use-case layer
type UserUsecase interface {
	Signup(entities.User) error
}

// Router represents router (multiplexer)
type Router interface {
	mapRoutes(UserHandler)
	http.Handler
}