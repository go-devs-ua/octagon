package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/ent"
)

type UserLogic interface {
	Signup(ent.User) error
}

type Router interface {
	mapRoutes(ah *UserHandler)
	http.Handler
}
