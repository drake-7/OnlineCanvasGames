package page

import (
	"net/http"
	"strconv"

	"github.com/1001bit/OnlineCanvasGames/internal/database"
	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
)

type ProfileData struct {
	Username string
	Date     string
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	data := ProfileData{}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HandleNotFound(w, r)
		return
	}

	user, err := usermodel.GetByID(id)
	if err != nil {
		HandleNotFound(w, r)
		return
	}

	data.Username = user.Name
	data.Date, err = database.FormatPostgresDate(user.Date)
	if err != nil {
		HandleServerError(w, r)
		return
	}

	serveTemplate("profile.html", data, w, r)
}
