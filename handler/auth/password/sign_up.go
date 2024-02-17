package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dev-xero/authentication-backend/authentication"
	"github.com/dev-xero/authentication-backend/model"
	"github.com/dev-xero/authentication-backend/service"
	"github.com/dev-xero/authentication-backend/util"
	validators "github.com/dev-xero/authentication-backend/validator"
	"github.com/google/uuid"
)

/*
Handles creating the user, assigning a JSON Web Token and responding with the user object

Objectives:
  - Decode and read the auth request body into a user object
  - Sanitize the user input
  - Validate the user input
  - Check that the user does not already exist
  - Create a JSON Web Token
  - Insert the user into the database
  - Respond with the user object payload

Params:
  - auth: The auth repo service
  - w:    A http response writer
  - r:    A pointer to a http request object

Returns:
  - No return value
*/
func SignUp(auth *service.AuthService, w http.ResponseWriter, r *http.Request) {
	// Store the request body
	var body = util.SignUpRequestBody{}

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
	userExists, err := auth.Repo.UserExists(r.Context(), body.Email)
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
	err = auth.Repo.InsertUser(r.Context(), user)
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
