package middleware

import (
	"log"
	"net/http"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authentication requested on:", r.URL)
		next.ServeHTTP(w, r)
	})
}
