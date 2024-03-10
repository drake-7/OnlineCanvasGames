package authapi

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

func AuthPage(w http.ResponseWriter, r *http.Request) {
	tmplloader.Templates.ExecuteTemplate(w, "auth.html", nil)
}
