package routes

import (
	"Sentitube/api/handlers"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/api/v1/analyze/youtube", handlers.PostVideoID)
	//r.Post("/api/analyze/instagram", handlers.PostInstagramPostID)
}
