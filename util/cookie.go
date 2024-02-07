package util

import (
	"net/http"
	"time"
)

func CreateTokenCookie(token string) http.Cookie {
	expirationTime := time.Now().Add(time.Hour).Unix()
	return http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
		Expires:  time.Unix(expirationTime, 0), // expires in an hour
	}
}

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
