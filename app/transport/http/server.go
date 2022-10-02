package http

import (
	"net/http"
)

type Server struct{ *http.Server }

func NewServer(r Router) *Server {
	r.Route()

	srv := new(http.Server)
	srv.Handler = r
	srv.Addr = ":8080"
	// TODO: Add some params

	return &Server{srv}
}

// Run will run our server
func (srv *Server) Run() error {
	return srv.ListenAndServe()
}
