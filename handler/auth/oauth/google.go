package handler

import (
	"net/http"

	"github.com/dev-xero/authentication-backend/service"
)

/*
Handles Google account sign-in with OAuth2.0

Objectives:
  - Setup auth config struct
  - Request authentication from Google
  - Handle auth callback

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignIn(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Google route hit"))
}

/*
Handles callbacks to Google sign-in

Objectives:
  - Handle auth callback
  - Save user to the database

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func GoogleSignInCallback(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Google callback route hit"))
}
