package handler

import (
	"fmt"
	"net/http"

	oauth "github.com/dev-xero/authentication-backend/handler/auth/oauth"
	password "github.com/dev-xero/authentication-backend/handler/auth/password"
	shared "github.com/dev-xero/authentication-backend/handler/auth/shared"
	"github.com/dev-xero/authentication-backend/service"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	service *service.AuthService
}

func (authHandler *AuthHandler) WithService(service *service.AuthService) {
	authHandler.service = service
}

/*
Handles requests made to the base auth route

Objectives:
  - Respond with an auth welcome message

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Auth route home"
	util.JsonResponse(w, msg, http.StatusOK, nil)
}

/*
Handles requests made to the auth/sign-up endpoint

Objectives:
  - Create a JSON Web Token cookie
  - Respond with the user as the payload

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	password.SignUp(auth.service, w, r)
}

/*
Handles requests made to the auth/sign-in route

Objectives:
  - Validate request body and sign-in user
  - Respond with the user as a payload

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	password.SignIn(auth.service, w, r)
}

/*
Handles requests made to the auth/oauth/google route

Objectives:
  - Handle the request
  - Return user data model on success

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	oauth.GoogleSignIn(auth.service, w, r)
}

/*
Handles callbacks from Google

Objectives:
  - Handle Google sign-in callbacks

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) GoogleSignInCallback(w http.ResponseWriter, r *http.Request) {
	oauth.GoogleSignInCallback(auth.service, w, r)
}

/*
Returns a generic failure response when oauth fails

Params:
  - w: A http response writer
  - r: The  http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) OAuthFailure(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "x")
	msg := fmt.Sprintf("Failed to sign-in using %s OAuth", util.CapitalizeFirstLetter(provider))
	util.JsonResponse(w, msg, http.StatusUnauthorized, nil)
}

/*
Handles requests made to the auth/sign-in route

Objectives:
  - Expire the cookie

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	shared.SignOut(w, r)
}
