package rest

import (
	"net/http"
	"regexp"
	"strings"

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
			logger.Infow("Request:",
				"Method", req.Method,
				"URL", req.URL,
				"User-Agent", req.UserAgent(),
			)
			h.ServeHTTP(w, req)
		})
	}
}

// WithValidateQuery check if query params is valid.
func WithValidateQuery(logger *lgr.Logger, allowedParams, allowedArgs *regexp.Regexp) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for param, arg := range req.URL.Query() {
				if !allowedParams.MatchString(param) {
					logger.Errorf("Unacceptable parameter %v in query", param)
					WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: "query params should match regex: " + allowedParams.String()}, logger)

					return
				}

				if !allowedArgs.MatchString(strings.Join(arg, "")) {
					logger.Errorf("Unacceptable argument in query parameter: %v", arg)
					WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: "query arguments should match regex: " + allowedArgs.String()}, logger)

					return
				}
			}
			h.ServeHTTP(w, req)
		})
	}
}
