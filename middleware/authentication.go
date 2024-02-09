package middleware

import (
	"log"
	"net/http"

	"github.com/dev-xero/authentication-backend/authentication"
	"github.com/dev-xero/authentication-backend/util"
)

/*
Authentication middleware for restricting access to protected routes

Objectives:
  - Obtain the token cookie if present
  - Verify the token
  - Authenticate the user based on whether the token is valid

Params:
  - next: A http handler

Returns:
  - A http handler
*/
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[LOG]: authentication requested on:", r.URL)

		tokenString, err := r.Cookie("token")
		if err != nil {
			log.Println("[FAIL]: token not present in cookie")
			msg := "Unauthorized request to a protected endpoint"
			util.JsonResponse(w, msg, http.StatusUnauthorized, nil)
			return
		}

		token, err := authentication.VerifyToken(tokenString.Value)
		if err != nil {
			log.Printf("[FAIL]: token verification failed: %v", err)
			msg := "Failed to verify token"
			util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
			return
		}

		log.Printf("[SUCCESS]: token successfully verified: %v", token.Claims)

		next.ServeHTTP(w, r)
	})
}
