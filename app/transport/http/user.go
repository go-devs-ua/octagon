package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user := entities.User{}

		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			return //someone have to work with errors
		}
		defer req.Body.Close()

		if err := user.Validate(); err != nil {
			return //someone have to work with errors
		}

		if err := uh.usecase.Signup(user); err != nil {
			return //someone have to work with errors
		}
	})
}
