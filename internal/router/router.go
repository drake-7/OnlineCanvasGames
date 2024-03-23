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

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", staticFileServer))

	// public
	router.Group(func(r chi.Router) {
		r.Post("/api/userauth", userauthapi.UserAuthPost)
	})

	// protected html
	router.Group(func(rAuth chi.Router) {
		rAuth.Use(AuthMiddlewareHTML)

		rAuth.Get("/", homeapi.HomePage)
		rAuth.Get("/game/{id}", gamepageapi.GamePage)
	})

	// protected non-Html
	router.Group(func(rAuth chi.Router) {
		rAuth.Use(AuthMiddleware)

		rAuth.Get("/some", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("data"))
		}))
	})

	return router
}
