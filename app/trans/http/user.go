package http

import (
	"net/http"
)

// addUser will create User
func (ah *ApiHandler) addUser(rw http.ResponseWriter, req *http.Request) {
	usr := User{}
	// some magic ...

	if err := usr.Validate(); err != nil {
		return
	}

	if err := ah.logic.Signup(usr); err != nil {
		return
	}

	// успіх
}
