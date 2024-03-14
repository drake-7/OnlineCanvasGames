package welcomeapi

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get only", http.StatusMethodNotAllowed)
		return
	}

	tmplloader.Templates.ExecuteTemplate(w, "welcome.html", nil)
}
