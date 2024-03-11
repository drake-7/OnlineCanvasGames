package router

import (
	"net/http"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Welcome
	mux.HandleFunc("/", welcomeapi.WelcomePage)
	mux.HandleFunc("/api/welcome", welcomeapi.WelcomePost)

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return mux
}
