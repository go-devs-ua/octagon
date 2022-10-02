package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mux represents router(multiplexer)
// and implements Router interface
type Mux struct {
	Router *mux.Router
	Map    []HandlerMap
}

// HandlerMap consists of all information
// that we need to register endpoints for specific entity
type HandlerMap struct {
	EndPoint string
	Handler  http.Handler
	Method   string
}

// NewRouter will initialise new instance of Mux
func NewRouter(hm ...HandlerMap) *Mux {
	return &Mux{
		Router: new(mux.Router),
		Map:    hm,
	}
}

// Route meth implements Router interface
// which makes Server initialisation router agnostic
// and allows us to switch to another router
// any time without massive refactoring
func (m *Mux) Route() {
	for _, hm := range m.Map {
		m.Router.Handle(hm.EndPoint, hm.Handler).Methods(hm.Method)
	}
}

// ServeHTTP is implementations of http.Handler interface
// that will handle HTTP (field Handler in http.Server)
func (m *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.Router.ServeHTTP(rw, req)
}
