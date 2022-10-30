package rest

import (
	"net/http"

	"github.com/go-devs-ua/octagon/lgr"
)

// Middleware is simple decorator.
type Middleware func(http.Handler, *lgr.Logger) http.Handler

// WrapMiddleware will build middleware chain.
func WrapMiddleware(h http.Handler, logger *lgr.Logger, middleware ...Middleware) http.Handler {
	for _, mw := range middleware {
		h = mw(h, logger)
	}

	return h
}

// WithLogRequest will log detailed request info.
func WithLogRequest(h http.Handler, logger *lgr.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Infow("Request:",
			"Method", req.Method,
			"URL", req.URL,
			"User-Agent", req.UserAgent(),
		)
		h.ServeHTTP(w, req)
	})
}
