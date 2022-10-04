package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(userHandler UserHandler) *Server {
	router := new(mux.Router)

	attachUserEndpoints(router, userHandler)

	return &Server{
		Server: &http.Server{
			Handler: router,
			Addr:    ":8080",
			// TODO: Add some params
		},
	}
}

// Run will run our server
func (s *Server) Run() error {
	return s.ListenAndServe()
}

func attachUserEndpoints(router *mux.Router, handler UserHandler) {
	router.Path("/users").Methods(http.MethodPost).HandlerFunc(handler.CreateUser)
	//router.Path("/users").Methods(http.MethodGet).HandlerFunc(handler.ListUsers)
}
