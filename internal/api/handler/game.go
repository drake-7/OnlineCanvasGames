package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type GameData struct {
	Game *gamemodel.Game
}

func GamePage(w http.ResponseWriter, r *http.Request) {
	data := GameData{
		Game: &gamemodel.Game{},
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		NotFound(w, r)
		return
	}

	data.Game.ID = id
	data.Game, err = gamemodel.ByID(id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			NotFound(w, r)
		default:
			ServerError(w, r)
		}
		return
	}

	tmplloader.ExecuteTemplate(w, r, "game.html", data)
}
