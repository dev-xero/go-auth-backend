package util

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

/*
Generates a random string and stores it as an oauth state cookie

Params:
  - w: A http response writer

Returns:
  - The oauth state string
*/
func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
