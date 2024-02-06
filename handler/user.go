package handler

import (
	"fmt"
	"net/http"

	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/go-chi/chi/v5"
)

type User struct {
	repo *repository.PostGreSQL
}

func (user *User) New(repo *repository.PostGreSQL) {
	user.repo = repo
}

func (user *User) Home(w http.ResponseWriter, r *http.Request) {
	msg := "User route home"
	util.JsonResponse(w, msg, http.StatusOK, nil)
}

func (user *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg := fmt.Sprintf("Get user by ID: %s hit", id)
	util.JsonResponse(w, msg, http.StatusOK, nil)
}
