package handler

import (
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
	fmt.Println("Sign-up route hit")
}

func (auth *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-in route hit")
}

func (auth *Auth) SignOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign-out route hit")
}
