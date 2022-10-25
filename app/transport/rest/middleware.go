package rest

import (
	"net/http"
	"regexp"
	"strings"
	"time"

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

// WithValidateQuery check if query params is valid.
func WithValidateQuery(allowedParams, allowedArgs *regexp.Regexp) Middleware {
	return func(h http.Handler, logger *lgr.Logger) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for param, arg := range req.URL.Query() {
				if !allowedParams.MatchString(param) {
					logger.Errorw("Unacceptable parameter.",
						MsgParam, param,
						MsgArg, arg,
					)
					WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: "query params should match regex: " + allowedParams.String()}, logger)

					return
				}

				if !allowedArgs.MatchString(strings.Join(arg, "")) {
					logger.Errorw("Unacceptable argument.",
						MsgParam, param,
						MsgArg, arg,
					)
					WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: "query arguments should match regex: " + allowedArgs.String()}, logger)

					return
				}
			}
			h.ServeHTTP(w, req)
		})
	}
}

// WithoutPanic will recover server from panic.
func WithoutPanic(h http.Handler, logger *lgr.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, logger)
				logger.Errorw(MsgPanic,
					MsgErr, err)
			}
		}()

		h.ServeHTTP(w, req)
	})
}

// WithHandlerTimeout set up handler timout.
func WithHandlerTimeout(h http.Handler, logger *lgr.Logger) http.Handler {
	return http.TimeoutHandler(h, handlerTimeoutSeconds*time.Second, MsgTimeOut)
}
