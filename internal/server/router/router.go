package router

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/api"
	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/page"
	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/socket"
	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/storage"
	"github.com/1001bit/OnlineCanvasGames/internal/server/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() (http.Handler, error) {
	router := chi.NewRouter()
	router.Use(chimw.Logger)
	router.Use(chimw.RedirectSlashes)

	// Storage
	router.Handle("/static/*", http.StripPrefix("/static", http.HandlerFunc(storage.HandleStatic)))
	router.Get("/favicon.ico", storage.HandleStatic)
	router.Handle("/image/*", http.StripPrefix("/image", http.HandlerFunc(storage.HandleImage)))

	// Websockets
	gamesWS, err := socket.NewGamesWS()
	if err != nil {
		return nil, err
	}
	go gamesWS.Run()
	router.HandleFunc("/ws/gameplay/{id}", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWS(gamesWS, w, r)
	})

	// API
	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.TypeJSON)

		// Post
		r.Post("/user", api.HandleUserPost)
	})

	// HTML Pages
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.TypeHTML)

		// Get
		r.Get("/", page.HandleHome)
		r.Get("/auth", page.HandleAuth)
		r.Get("/profile/{id}", page.HandleProfile)
		r.Get("/game/{id}", page.HandleGame)
		r.Group(func(sr chi.Router) {
			sr.Use(middleware.AuthHTML)
			sr.Get("/game/{id}/play", page.HandleGameplay)
		})
		r.Get("/*", page.HandleNotFound)
	})

	return router, nil
}
