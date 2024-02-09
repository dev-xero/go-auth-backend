package util

import (
	"net/http"
	"time"
)

/*
Creates a cookie with the token, with a max age of 1 hour

Objectives:
  - Create a cookie with the value set to the token to expire in 1 hour

Params:
  - token: A JSON Web Token to create the cookie with

Returns:
  - A http cookie with the token and configurations
*/
func CreateTokenCookie(token string) http.Cookie {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
		MaxAge:   3600, // Lives for 1 hour
	}
	return cookie
}

/*
Expires any cookies saved in the client

Objectives:
  - Set the max age property of the cookie to < 0

Params:
  - w:    A http response writer
  - name: The name of the cookie to expire

Returns:
  - No return value
*/
func ExpireCookie(w http.ResponseWriter, name string) {
	expiration := time.Now().Add(-24 * time.Hour)
	deletedCookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Expires:  expiration,
		MaxAge:   -1,
		HttpOnly: false,
	}
	http.SetCookie(w, deletedCookie)
}
