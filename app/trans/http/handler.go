package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/ent"
)

type User = ent.User

// ApiHandler is User HTTP handler
type ApiHandler struct {
	logic UserLogic
}

// NewApiHandler will return *ApiHandler
// accepting UseCases interface
func NewApiHandler(logic UserLogic) *ApiHandler {
	return &ApiHandler{
		logic: logic,
	}
}

// HandleUser will take care of our handsome User's endpoint
func (ah *ApiHandler) HandleUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		// POST ...
		}
	})
}
