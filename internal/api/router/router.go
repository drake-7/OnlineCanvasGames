package router

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/api/handler"
	"github.com/1001bit/OnlineCanvasGames/internal/api/middleware"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(chimw.Logger)
	router.Use(chimw.RedirectSlashes)

	// Storage
	router.Handle("/static/*", http.StripPrefix("/static", http.HandlerFunc(handler.StaticStorage)))
	router.Handle("/favicon.ico", http.HandlerFunc(handler.StaticStorage))
	router.Handle("/image/*", http.StripPrefix("/image", http.HandlerFunc(handler.ImageStorage)))

	router.Group(func(r chi.Router) {
		r.Use(middleware.TypeJSON)

		// Post
		router.Post("/api/user", handler.UserPost)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.TypeHTML)

		// Get
		router.Get("/", handler.HomePage)
		router.Get("/auth", handler.AuthPage)
		router.Get("/game/{id}", handler.GamePage)
		router.Get("/profile/{id}", handler.ProfilePage)

	})

	return router
}
