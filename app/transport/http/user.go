package http

import (
	"fmt"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

func (uh UserHandler) Map() []HandlerMap {
	return []HandlerMap{
		{EndPoint: "/user", Handler: uh.CreateUser(), Method: http.MethodGet},
		// {EndPoint: "/user", Handler: uh.DeleteUser(), Method: http.MethodDelete},
		// and so on...
	}
}

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		usr := entities.User{}

		if err := usr.Validate(); err != nil {
			return
		}
		fmt.Fprintf(rw, "bla bla %d", 5)
		if err := uh.usecase.Signup(usr); err != nil {
			return
		}
	})
}
