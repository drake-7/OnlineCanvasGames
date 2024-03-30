package page

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	// only unauthorized may see this page
	_, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	serveTemplate("auth.html", nil, w, r)
}
