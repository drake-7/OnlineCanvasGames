package router

import (
	"net/http"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: JWT
		if true {
			welcomeapi.WelcomePage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
