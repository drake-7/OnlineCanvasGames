package handler

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	tmplloader.Templates.ExecuteTemplate(w, "auth.html", nil)
}

func ServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusUnauthorized)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusUnauthorized)
}
