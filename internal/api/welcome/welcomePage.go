package welcomeapi

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	tmplloader.Templates.ExecuteTemplate(w, "welcome.html", nil)
}
