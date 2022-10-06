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

// Logger represents logger
type Logger interface {
	Errorf(string, ...any)
	Debugf(string, ...any)
	Infof(string, ...any)
	Warnf(string, ...any)
}
