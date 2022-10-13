package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/lgr"
)

// Middleware is simple decorator
type Middleware func(http.Handler) http.Handler

// WrapMiddleware will build middleware chain
func WrapMiddleware(handler http.Handler, middleware ...Middleware) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}

// WithLogRequest will log detailed request info
func WithLogRequest(logger *lgr.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer logger.Infow("Request:",
				"Header", req.Header,
				"Method", req.Method,
				"URL", req.URL,
				"User-Agent", req.UserAgent(),
			)
			h.ServeHTTP(w, req)
		})
	}
}
