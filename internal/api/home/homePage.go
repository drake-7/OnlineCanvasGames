package homeapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type HomeData struct {
	Name  string
	Games []gamemodel.Game
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	data := HomeData{}

	// get name from jwt
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := auth.GetJWTClaims(cookie.Value)
	if err != nil {
		http.Error(w, "bad token", http.StatusUnauthorized)
	}

	data.Name = fmt.Sprint(claims["username"])

	// games count
	data.Games, err = gamemodel.All()
	if err != nil {
		data.Games = nil
		log.Println("error getting games:", err)
	}

	tmplloader.Templates.ExecuteTemplate(w, "home.html", data)
}
