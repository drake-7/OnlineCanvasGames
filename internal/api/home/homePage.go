package homeapi

import (
	"fmt"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type Game struct {
	Title string
}

type HomeData struct {
	Name  string
	Games []Game
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get only", http.StatusMethodNotAllowed)
		return
	}

	data := HomeData{}

	// get name from jwt
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

	// games count
	data.Games = append(data.Games, Game{Title: "funny game"})

	tmplloader.Templates.ExecuteTemplate(w, "home.html", data)
}
