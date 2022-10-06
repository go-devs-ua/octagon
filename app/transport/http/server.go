package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-devs-ua/octagon/cfg"
	"github.com/gorilla/mux"
)

const timeoutMsg = "Connection timeout"

// Server is simple server
type Server struct{ *http.Server }

type Handlers struct {
	UserHandler UserHandler
}

// NewServer will initialize the server
func NewServer(opt cfg.Options, handlers Handlers) *Server {
	router := new(mux.Router)

	attachUserEndpoints(router, handlers)

	return &Server{
		Server: &http.Server{
			Addr:         opt.Server.Host + ":" + opt.Server.Port,
			Handler:      http.TimeoutHandler(router, 3*time.Second, timeoutMsg),
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

// Run will run our server
func (srv *Server) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}

	return nil
}

func attachUserEndpoints(router *mux.Router, handlers Handlers) {
	router.Path("/users").Methods(http.MethodPost).Handler(handlers.UserHandler.CreateUser())
}
