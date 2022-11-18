package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/gorilla/mux"
)

// Server is simple server.
type Server struct{ *http.Server }

type Handlers struct {
	UserHandler UserHandler
}

// NewServer will initialize the server.
func NewServer(opt cfg.Options, handlers Handlers, logger *lgr.Logger) *Server {
	router := new(mux.Router)

	attachUserEndpoints(router, handlers)
	handler := WrapMiddlewares(router, logger, WithLogRequest)

	return &Server{
		Server: &http.Server{
			Addr:         opt.Server.Host + ":" + opt.Server.Port,
			Handler:      handler,
			ReadTimeout:  readTimeoutSeconds * time.Second,
			WriteTimeout: writeTimeoutSeconds * time.Second,
		},
	}
}

// Run will run our server.
func (srv *Server) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}

	return nil
}

func attachUserEndpoints(router *mux.Router, handlers Handlers) {
	router.Path("/users").Methods(http.MethodPost).HandlerFunc(handlers.UserHandler.CreateUser)
	router.Path("/users").Methods(http.MethodGet).HandlerFunc(handlers.UserHandler.GetAllUsers)
	router.Path("/users/{id}").Methods(http.MethodGet).HandlerFunc(handlers.UserHandler.GetUserByID)
	router.Path("/users").Methods(http.MethodDelete).HandlerFunc(handlers.UserHandler.DeleteUser)
}
