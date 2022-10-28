package rest

import (
	"net/http"

	"github.com/go-devs-ua/octagon/lgr"
)

// Middleware is simple decorator.
type Middleware func(http.Handler) http.Handler

// WrapMiddleware will build middleware chain.
func WrapMiddleware(h http.Handler, middleware ...Middleware) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}

	return h
}

// WithLogRequest will log detailed request info.
func WithLogRequest(logger *lgr.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logger.Debugw("Request:", "Method", req.Method, "URL", req.URL, "User-Agent", req.UserAgent())
			h.ServeHTTP(w, req)
		})
	}
}
