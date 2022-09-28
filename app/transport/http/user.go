package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/ent"
)

type User = ent.User

// Handle will take care of our handsome User's endpoint
func (uh *UserHandler) Handle() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			uh.add(rw, req)
		default:
		}
	})
}

// add will create User
func (uh *UserHandler) add(rw http.ResponseWriter, req *http.Request) {
	usr := User{}
	// some magic ...

	if err := usr.Validate(); err != nil {
		return
	}

	if err := uh.logic.Signup(usr); err != nil {
		return
	}

	// успіх
}
