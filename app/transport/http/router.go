package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mux is simple router
// that implements Router interface.
type Mux struct{ mux.Router }

func NewRouter() *Mux {
	return new(Mux)
}

// mapRoutes will take care of all endpoints
func (m *Mux) mapRoutes(ah UserHandler) {
	m.Handle("/user", ah.CreateUser()).Methods(http.MethodPost)
}
