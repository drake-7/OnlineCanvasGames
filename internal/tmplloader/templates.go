package tmplloader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

type HeaderData struct {
	Username string
	UserID   string
}

type TemplateData struct {
	Data   any
	Header HeaderData
}

var templates = template.Must(template.ParseGlob("web/templates/**/*.html"))

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, file string, data any) {
	tmplData := TemplateData{
		Data: data,
	}

	claims, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		tmplData.Header.UserID = fmt.Sprint(claims["userID"])
		tmplData.Header.Username = fmt.Sprint(claims["username"])
	}

	templates.ExecuteTemplate(w, file, tmplData)
}
