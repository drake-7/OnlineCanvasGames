package router

import (
	"net/http"
	"time"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token from cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			welcomeapi.WelcomePage(w, r)
			return
		}

		tokenString := cookie.Value

		// get token claims
		claims, err := auth.GetJWTClaims(tokenString)
		if err != nil {
			welcomeapi.WelcomePage(w, r)
			return
		}

		// check token expiry
		expTime, err := claims.GetExpirationTime()
		if err != nil || time.Now().Unix() > expTime.Unix() {
			welcomeapi.WelcomePage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
