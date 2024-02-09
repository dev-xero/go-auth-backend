package util

import (
	"net/http"
	"time"
)

func CreateTokenCookie(token string) http.Cookie {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
		MaxAge:   60, // Lives for 1 hour
	}
	return cookie
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
