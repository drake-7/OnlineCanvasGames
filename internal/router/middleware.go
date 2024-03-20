package router

import (
	"net/http"
	"time"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

// check JWT for any handler
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token from cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// get token claims
		claims, err := auth.GetJWTClaims(cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// check token expiry
		expTime, err := claims.GetExpirationTime()
		if err != nil || time.Now().Unix() > expTime.Unix() {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// check JWT for HTML handlers
func AuthMiddlewareHTML(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token from cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			welcomeapi.WelcomePage(w, r)
			return
		}

		// get token claims
		claims, err := auth.GetJWTClaims(cookie.Value)
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
