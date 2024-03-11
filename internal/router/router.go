package router

import (
	"net/http"

	userauthapi "github.com/1001bit/OnlineCanvasGames/internal/api/userauth"
	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Welcome
	mux.HandleFunc("/", welcomeapi.WelcomePage)
	mux.HandleFunc("/api/userauth", userauthapi.UserAuthPost)

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return mux
}
