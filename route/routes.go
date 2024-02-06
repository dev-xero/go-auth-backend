package route

import (
	"database/sql"
	"net/http"

	auth "github.com/dev-xero/authentication-backend/route/auth"
	user "github.com/dev-xero/authentication-backend/route/user"
	"github.com/dev-xero/authentication-backend/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Load routes in the app server
func LoadRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Setup cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Welcome to the API"
		util.JsonResponse(w, msg, http.StatusOK, nil)
	})

	// Setup auth route handlers
	router.Route("/auth", func(router chi.Router) {
		auth.LoadAuthRoutes(router, db)
	})

	// Setup user route handlers
	router.Route("/user", func(router chi.Router) {
		user.LoadUserRoutes(router, db)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		msg := "Undefined endpoint accessed"
		util.JsonResponse(w, msg, http.StatusNotFound, nil)
	})

	return router
}
