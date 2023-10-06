package api

import (
	"Sentitube/api/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Start() {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize routes
	routes.InitializeRoutes(r)

	// Start the server
	http.ListenAndServe(":8080", r)
}
