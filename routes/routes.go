package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// All returns a router containing the api routes.
func All() chi.Router {
	r := chi.NewRouter()

	// setup CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Cache-Control"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(crs.Handler)

	r.Route("/api/", func(r chi.Router) {
		r.Mount("/planets", planets())
	})

	return r
}
