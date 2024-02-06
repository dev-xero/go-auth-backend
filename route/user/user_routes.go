package route

import (
	"database/sql"

	"github.com/dev-xero/authentication-backend/handler"
	repository "github.com/dev-xero/authentication-backend/repository/user"
	"github.com/go-chi/chi/v5"
)

func LoadUserRoutes(router chi.Router, db *sql.DB) {
	user := &handler.User{}
	user.New(&repository.PostGreSQL{Database: db})

	router.Get("/", user.Home)
	router.Get("/{id}", user.GetUserByID)
}
