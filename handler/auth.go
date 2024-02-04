package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dev-xero/authentication-backend/util"
)

type Auth struct{}

func (auth *Auth) Home(w http.ResponseWriter, r *http.Request) {
	msg := "Auth route home"
	util.JsonResponse(w, msg, http.StatusOK, nil)
}

func (auth *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Read response body into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		msg := "Bad request, email or password not present"
		util.JsonResponse(w, msg, http.StatusBadRequest, nil)
		return
	}

	fmt.Printf("Email: %s\nPassword: %s", body.Email, body.Password)
}

func (auth *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-in route hit")
}

func (auth *Auth) SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-out route hit")
}
