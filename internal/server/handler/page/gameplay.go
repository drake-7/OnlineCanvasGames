package page

import (
	"database/sql"
	"net/http"
	"strconv"

	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
)

func HandleGameplay(w http.ResponseWriter, r *http.Request) {
	data := GameData{
		Game: &gamemodel.Game{},
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HandleNotFound(w, r)
		return
	}

	data.Game.ID = id
	data.Game, err = gamemodel.GetByID(id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			HandleNotFound(w, r)
		default:
			HandleServerError(w, r)
		}
		return
	}

	serveTemplate("gameplay.html", data, w, r)
}
