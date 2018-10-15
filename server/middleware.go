package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type logger interface {
	Println(v ...interface{})
}

func loggingMiddleware(logger logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			logger.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
