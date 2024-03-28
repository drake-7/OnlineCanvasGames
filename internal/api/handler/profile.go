package handler

import (
	"net/http"
	"strconv"

	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type ProfileData struct {
	Username string
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

	tmplloader.ExecuteTemplate(w, r, "profile.html", data)
}
