package router

import (
	"net/http"

	homeapi "github.com/1001bit/OnlineCanvasGames/internal/api/home"
	userauthapi "github.com/1001bit/OnlineCanvasGames/internal/api/userauth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// userauth api
	mux.HandleFunc("/api/userauth", userauthapi.UserAuthPost)

	// home
	mux.HandleFunc("/", AuthMiddlewareHTML(homeapi.HomePage))

	// TEST CASE
	mux.Handle("/some", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("data"))
	})))

	// static
	staticFileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return mux
}
