package page

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
)

type HomeData struct {
	Username string
	Games    []gamemodel.Game
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	data := HomeData{}

	claims, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		data.Username = fmt.Sprint(claims["username"])
	}

	// games count
	data.Games, err = gamemodel.GetAll()
	if err != nil {
		data.Games = nil
		log.Println("error getting games:", err)
	}

	serveTemplate("home.html", data, w, r)
}
