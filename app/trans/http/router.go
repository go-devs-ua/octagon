package http

import (
	"net/http"
)

// Mux is simple router
// that implements Router interface.
type Mux struct{ http.ServeMux }

func NewRouter() *Mux {
	return new(Mux)
}

func (mux *Mux) mapRoutes(ah *ApiHandler) {
	mux.Handle("/user", ah.HandleUser())
}
