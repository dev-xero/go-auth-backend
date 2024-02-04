package route

import (
	"database/sql"
	"net/http"

	route "github.com/dev-xero/authentication-backend/route/auth"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Load routes in the app server
func LoadRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Welcome to the API"
		util.JsonResponse(w, msg, http.StatusOK, nil)
	})

	// Setup auth route handlers
	router.Route("/auth", func(router chi.Router) {
		route.LoadAuthRoutes(router, db)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		msg := "Undefined endpoint accessed"
		util.JsonResponse(w, msg, http.StatusNotFound, nil)
	})

	return router
}
