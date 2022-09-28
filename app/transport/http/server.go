package http

import (
	"fmt"
	"net/http"
)

// Server is simple server
type Server struct{ *http.Server }

// NewServer will initialize the server
// that would be router type agnostic
// we can switch to any router type
// that implements Router interface
func NewServer(uu UserUsecase, r Router) *Server {
	hdl := NewUserHandler(uu)
	r.mapRoutes(hdl)
	// TODO: Add config
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &Server{srv}
}

// Run will run our server
func (srv *Server) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}

	return nil
}
