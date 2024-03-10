package router

import (
	"net/http"

	authapi "github.com/1001bit/OnlineCanvasGames/internal/api/auth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", authapi.AuthPage)

	return mux
}
