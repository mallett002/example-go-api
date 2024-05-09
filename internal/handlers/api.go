package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/mallett002/example-go-api/internal/middleware"
)

// capital "Handler" can be used in other files (public function)
func Handler(router *chi.Mux) {
	// Global middleware (applied all the time)
	router.Use(chimiddle.StripSlashes) // removes tailing slashes (can give 404s)

	router.Route("/account", func(router chi.Router) {

		// middleware for authorization
		router.Use(middleware.Authorization) // /internal/middleware

		router.Get("/coins", GetCoinBalance)
	})
}
