package router

import (
	"net/http"

	gamepageapi "github.com/1001bit/OnlineCanvasGames/internal/api/gamepage"
	homeapi "github.com/1001bit/OnlineCanvasGames/internal/api/home"
	userauthapi "github.com/1001bit/OnlineCanvasGames/internal/api/userauth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RedirectSlashes)

	// Non-Protected Routes
	// Non-html
	router.Post("/api/userauth", userauthapi.UserAuthPost)

	// Protected Routes
	// Html
	router.Group(func(r chi.Router) {
		r.Use(AuthMiddlewareHTML)

		r.Get("/", homeapi.HomePage)
		r.Get("/game/{id}", gamepageapi.GamePage)
	})

	// Non-Html
	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)

		// TEST CASE
		r.Get("/some", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("data"))
		}))
	})

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", staticFileServer))

	return router
}
