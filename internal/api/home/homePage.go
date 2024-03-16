package homeapi

import (
	"fmt"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type HomeData struct {
	Name string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get only", http.StatusMethodNotAllowed)
		return
	}

	data := HomeData{}

	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := auth.GetJWTClaims(&cookie.Value)
	if err != nil {
		http.Error(w, "bad token", http.StatusUnauthorized)
	}

	data.Name = fmt.Sprint(claims["username"])

	tmplloader.Templates.ExecuteTemplate(w, "home.html", data)
}
