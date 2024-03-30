package handler

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

func AuthPage(w http.ResponseWriter, r *http.Request) {
	// only unauthorized may see this page
	_, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmplloader.ExecuteTemplate(w, r, "auth.html", nil)
}
