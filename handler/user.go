package handler

import (
	"fmt"
	"net/http"
	"strings"

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
	var msg = ""

	// Get user from the database
	theUser, err := user.repo.GetUserByID(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			msg = "A user with that id doesn't exist"
			util.JsonResponse(w, msg, http.StatusBadRequest, nil)
			return
		}
		msg = "Internal server error, failed to get user with that id"
		util.JsonResponse(w, msg, http.StatusInternalServerError, nil)
		return
	}

	msg = fmt.Sprintf("Successfully fetched user with the id: %s", id)
	util.JsonResponse(w, msg, http.StatusOK, theUser)
}
