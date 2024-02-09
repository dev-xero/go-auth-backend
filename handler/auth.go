package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev-xero/authentication-backend/authentication"
	"github.com/dev-xero/authentication-backend/model"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/dev-xero/authentication-backend/util"
	validators "github.com/dev-xero/authentication-backend/validator"
	"github.com/google/uuid"
)

/*
Auth handler struct

Objectives:
  - Handle all auth requests

Fields:
  - repo: The database repository
*/
type Auth struct {
	repo *repository.PostGreSQL
}

/*
Initializes a new auth request handler

Objectives:
  - Initialize an auth request handler with the provided repo

Params:
  - repo: The database repo to bind the handler to

Returns:
  - No return value
*/
func (auth *Auth) New(repo *repository.PostGreSQL) {
	auth.repo = repo
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
func (auth *Auth) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Auth route home"
	util.JsonResponse(w, msg, http.StatusOK, nil)
}

/*
Handles requests made to the auth/sign-up endpoint

Objectives:
  - Decode and read the auth request body into a user object
  - Sanitize the user input
  - Validate the user input
  - Check that the user does not already exist
  - Create a JSON Web Token
  - Insert the user into the database
  - Respond with the user object payload

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	// Store the request body
	var body = util.AuthRequestBody{}

	// Read response body into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		msg := "Bad request, username, email or password not present"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Sanitize user input
	util.SanitizeUserInput(&body)
	log.Printf("[LOG]: Sanitized user input")

	// Validate the user input
	if err := validators.ValidateUserInput(&body); err != nil {
		msg := util.CapitalizeFirstLetter(err.Error())
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Do not create a new user if the user already exists
	userExists, err := auth.repo.UserExists(r.Context(), body.Email)
	if err != nil {
		log.Println(err)
		msg := "Could not check if user already exists"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Respond with a bad request if the user exists
	if userExists {
		msg := "A user with that email already exists"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Prepare the user data for insertion
	var user = model.User{
		ID:       uuid.New(),
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}

	// Generate a JSON Web token that can be sent to the user
	token, err := authentication.CreateJWToken(user.ID)
	if err != nil {
		log.Println(err)
		msg := "Failed to create token"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Insert the user into the database
	err = auth.repo.InsertUser(r.Context(), user)
	if err != nil {
		log.Println(err)
		msg := "Could not insert user into database"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Create the user payload
	var userPayload = util.UserPayload{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Set the token cookie and send the response
	cookie := util.CreateTokenCookie(token)
	http.SetCookie(w, &cookie)
	util.JsonResponse(w, "Successfully inserted user into database", http.StatusOK, userPayload)
}

/*
Handles requests made to the auth/sign-in route

Objectives:
  - Decode and read the auth request body into a user object
  - Expire any token cookies that may be present
  - Sanitize the user input
  - Check that the user  exists
  - If the use does not exist, respond with an error
  - Compare the request body password with the user password hash
  - Generate a new JSON Web token
  - Send the token cookie and the user payload response

Params:
  - w: A http response writer
  - r: A pointer to a http request object

Returns:
  - No return value
*/
func (auth *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	// Store the auth request body
	var body = util.AuthRequestBody{}

	// Read response body into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		msg := "Bad request, username, email or password not present"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Expire the token cookie
	util.ExpireCookie(w, "token")

	// Sanitize user input
	util.SanitizeUserInput(&body)

	// Check if the user exists
	userExists, err := auth.repo.UserExists(r.Context(), body.Email)
	if err != nil {
		msg := "Internal server error, could not check if user already exists"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// If the user does not exist, respond with an error
	if !userExists {
		msg := "A user with those credentials does not exist"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Query the database and obtain the user the the provided email
	user, err := auth.repo.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		msg := "Internal server error, could not check if a user with that email exists"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Check that the password matches the hash
	if !util.CompareWithHash([]byte(user.Password), body.Password) {
		msg := "Provided passwords mismatch"
		util.JsonResponse(w, msg, http.StatusUnauthorized, nil)
		return
	}

	// Generate a new token and send response
	token, err := authentication.CreateJWToken(user.ID)
	if err != nil {
		log.Println(err)
		msg := "Failed to create token"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Set the token cookie and send the response
	cookie := util.CreateTokenCookie(token)
	http.SetCookie(w, &cookie)
	util.JsonResponse(w, "Successfully signed-in", http.StatusOK, user)
}

func (auth *Auth) SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-out route hit")
}
