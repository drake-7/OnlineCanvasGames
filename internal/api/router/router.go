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

	// PUBLIC
	// Static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", staticFileServer))

	// Storage
	storageFileServer := http.FileServer(http.Dir("./web/storage/"))
	router.Handle("/storage/*", http.StripPrefix("/storage", storageFileServer))

	// Plaintext
	router.Group(func(r chi.Router) {
		r.Post("/api/userauth", handler.AuthPost)
	})

	// PROTECTED
	// Plaintext
	router.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Get("/some", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("data"))
		}))
	})

	// HTML
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthHTML)

		r.Get("/", handler.HomePage)
		r.Get("/game/{id}", handler.GamePage)
	})

	return router
}
