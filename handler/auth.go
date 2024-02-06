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
	"github.com/mrz1836/go-sanitize"
)

type Auth struct {
	repo *repository.PostGreSQL
}

func (auth *Auth) New(repo *repository.PostGreSQL) {
	auth.repo = repo
}

func (auth *Auth) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Auth route home"
	util.JsonResponse(w, msg, http.StatusOK, nil)
}

func (auth *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	var body = util.AuthRequestBody{}

	// Read response body into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		msg := "Bad request, username, email or password not present"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	// Sanitize user input
	body.Email = sanitize.Email(body.Email, false)
	body.Username = sanitize.Alpha(body.Username, false)
	body.Password = sanitize.AlphaNumeric(body.Password, false)

	log.Printf("[LOG]: Sanitized user input")

	if err := validators.ValidateUserInput(&body); err != nil {
		msg := util.CapitalizeFirstLetter(err.Error())
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	userExists, err := auth.repo.UserExists(r.Context(), body.Email)
	if err != nil {
		log.Println(err)
		msg := "Could not check if user already exists"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	// Check that the user doesn't already exist
	if userExists {
		msg := "A user with that email already exists"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	var user = model.User{
		ID:       uuid.New(),
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}

	// Create a token and send it back to the user
	token, err := authentication.CreateJWToken(user.ID)
	if err != nil {
		log.Println(err)
		msg := "Failed to create token"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	err = auth.repo.InsertUser(r.Context(), user)
	if err != nil {
		log.Println(err)
		msg := "Could not insert user into database"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
	}

	http.SetCookie(w, &cookie)

	// Return user object in response
	msg := "Successfully inserted user into database"
	util.JsonResponse(w, msg, http.StatusOK, user)
}

func (auth *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-in route hit")
}

func (auth *Auth) SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-out route hit")
}
