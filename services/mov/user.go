package mov

import (
	"net/http"

	"github.com/go-devs-ua/octagon/core"
)

type User = core.User

// HandleUser will take care of our handsome User's endpoint
func (mvr *Mover) HandleUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	})
}

// Add will create User
func (mvr *Mover) addUser(rw http.ResponseWriter, req *http.Request) {}
