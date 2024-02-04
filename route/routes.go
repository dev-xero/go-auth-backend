package route

import (
	"net/http"

	"github.com/dev-xero/authentication-backend/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Load routes in the app server
func LoadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Welcome to the API"
		util.JsonResponse(w, msg, http.StatusOK, nil)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		msg := "Undefined endpoint accessed"
		util.JsonResponse(w, msg, http.StatusNotFound, nil)
	})

	return router
}
