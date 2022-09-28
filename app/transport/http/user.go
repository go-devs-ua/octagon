package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		usr := entities.User{}

		if err := usr.Validate(); err != nil {
			return
		}

		if err := uh.usecase.Signup(usr); err != nil {
			return
		}
	})
}
