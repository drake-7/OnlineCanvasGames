package router

import (
	"net/http"

	authapi "github.com/1001bit/OnlineCanvasGames/internal/api/auth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Auth
	mux.HandleFunc("/auth", authapi.AuthPage)

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return mux
}
