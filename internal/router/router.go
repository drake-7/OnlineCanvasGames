package router

import (
	"net/http"

	authapi "github.com/1001bit/OnlineCanvasGames/internal/api/auth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth", authapi.AuthPage)

	return mux
}
