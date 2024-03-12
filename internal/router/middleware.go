package router

import (
	"log"
	"net/http"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			switch {
			case err == http.ErrNoCookie:
				welcomeapi.WelcomePage(w, r)
			default:
				log.Println(err)
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}

		token := cookie.Value

		err = auth.VerifyJWT(token)
		if err != nil {
			welcomeapi.WelcomePage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
