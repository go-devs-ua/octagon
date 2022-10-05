package http

import (
	"fmt"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const timeoutMsg = "Connection timeout"

// Server is simple server
type Server struct{ *http.Server }

// NewServer will initialize the server
func NewServer(opt cfg.Options, handler UserHandler) *Server {
	router := new(mux.Router)

	attachUserEndpoints(router, handler)

	return &Server{
		Server: &http.Server{
			Addr:         opt.Server.Host,
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

func attachUserEndpoints(router *mux.Router, handler UserHandler) {
	router.Path("/users").Methods(http.MethodPost).Handler(handler.CreateUser())
}
