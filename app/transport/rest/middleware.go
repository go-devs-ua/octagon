package rest

import (
	"net/http"

	"github.com/go-devs-ua/octagon/lgr"
)

// Middleware is simple decorator.
type Middleware func(http.Handler, *lgr.Logger) http.Handler

// WrapMiddlewares will build middleware chain.
func WrapMiddlewares(h http.Handler, logger *lgr.Logger, middlewares ...Middleware) http.Handler {
	for _, mw := range middlewares {
		h = mw(h, logger)
	}

	return h
}

// WithLogRequest will log detailed request info.
func WithLogRequest(h http.Handler, logger *lgr.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Debugw("Request:",
			"Method", req.Method,
			"URL", req.URL,
			"User-Agent", req.UserAgent(),
		)
		h.ServeHTTP(w, req)
	})
}
