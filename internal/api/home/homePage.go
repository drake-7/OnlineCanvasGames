package homeapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/model"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type HomeData struct {
	Name  string
	Games []model.Game
}

func getGames() ([]model.Game, error) {
	stmt, err := database.DB.GetStatement("getGames")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []model.Game

	for rows.Next() {
		var game model.Game

		err := rows.Scan(&game.ID, &game.Title)
		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
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
	data.Games, err = getGames()
	if err != nil {
		data.Games = nil
		log.Println("error getting games:", err)
	}

	tmplloader.Templates.ExecuteTemplate(w, "home.html", data)
}
