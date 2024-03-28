package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type HomeData struct {
	Username string
	Games    []gamemodel.Game
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	data := HomeData{}

	claims, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		data.Username = fmt.Sprint(claims["username"])
	}

	// games count
	data.Games, err = gamemodel.All()
	if err != nil {
		data.Games = nil
		log.Println("error getting games:", err)
	}

	tmplloader.ExecuteTemplate(w, r, "home.html", data)
}
