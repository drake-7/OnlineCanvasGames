package handler

import (
	"net/http"
	"strconv"

	"github.com/1001bit/OnlineCanvasGames/internal/database"
	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type ProfileData struct {
	Username string
	Date     string
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	data := ProfileData{}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		NotFound(w, r)
		return
	}

	user, err := usermodel.ByID(id)
	if err != nil {
		NotFound(w, r)
		return
	}

	data.Username = user.Name
	data.Date, err = database.FormatPostgresDate(user.Date)
	if err != nil {
		ServerError(w, r)
		return
	}

	tmplloader.ExecuteTemplate(w, r, "profile.html", data)
}
