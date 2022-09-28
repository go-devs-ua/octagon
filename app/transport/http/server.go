package http

import (
	"fmt"
	"net/http"
)

// Server is simple server
type Server struct {
	*http.Server
}

// NewServer will initialize the server
// that would be router type agnostic
// we can switch to any router type
// that implements Router interface
func NewServer(logic UserLogic, mux Router) *Server {
	srv := new(http.Server)
	hdl := NewApiHandler(logic)
	mux.mapRoutes(hdl)
	// TODO: Add config and stuff
	srv.Handler = mux
	srv.Addr = ":8080"
	return &Server{srv}
}

// Run will run our server
func (srv Server) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}

	return nil
}
