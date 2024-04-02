package middleware

import (
	"errors"
	"net/http"
)

var NoAuthError = errors.New("Not authed.")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do later I'm lazy as fuck
		next.ServeHTTP(w, r)
	})
}
