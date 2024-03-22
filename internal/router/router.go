package router

import (
	"net/http"

	gamepageapi "github.com/1001bit/OnlineCanvasGames/internal/api/gamepage"
	homeapi "github.com/1001bit/OnlineCanvasGames/internal/api/home"
	userauthapi "github.com/1001bit/OnlineCanvasGames/internal/api/userauth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// userauth api
	mux.HandleFunc("POST /api/userauth", userauthapi.UserAuthPost)

	// home
	mux.HandleFunc("GET /", AuthMiddlewareHTML(homeapi.HomePage))

	// game page
	mux.HandleFunc("GET /game/{id}", AuthMiddlewareHTML(gamepageapi.GamePage))

	// TEST CASE
	mux.Handle("GET /some", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("data"))
	})))

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", staticFileServer))

	return mux
}
