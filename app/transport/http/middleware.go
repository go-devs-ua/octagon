package http

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func WithLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(w, req)
		log.Println("BALLLLLLLLLLLLLLLLL")
	})
}
