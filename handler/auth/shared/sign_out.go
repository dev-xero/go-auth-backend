package handler

import (
	"log"
	"net/http"

	"github.com/dev-xero/authentication-backend/service"
	"github.com/dev-xero/authentication-backend/util"
)

/*
Handles signing-out the user and expiring tokens

Objectives:
  - Expire the token cookie
  - Redirect the user to the home route

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func SignOut(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	// Expire the token cookie
	util.ExpireCookie(w, "token")

	util.JsonResponse(w, "Successfully signed-out", http.StatusOK, nil)
	log.Println("[LOG]: Successfully signed user out")
}
