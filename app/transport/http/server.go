package http

import (
	"fmt"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Server is simple server
type Server struct{ *http.Server }

// NewServer will initialize the server
func NewServer(opt cfg.Options, uc UserUsecase) *Server {
	r := mux.NewRouter()

	h := NewUserHandler(uc)
	r.Handle("/user", h.CreateUser()).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         opt.Server.Host,
		Handler:      http.TimeoutHandler(r, 3*time.Second, "Connection timeout"),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
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
