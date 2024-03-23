package router

import (
	"net/http"
	"time"

	welcomeapi "github.com/1001bit/OnlineCanvasGames/internal/api/welcome"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

type Middleware func(next http.Handler) http.Handler

func checkAuth(r *http.Request) bool {
	// get token from cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return false
	}

	// get token claims
	claims, err := auth.GetJWTClaims(cookie.Value)
	if err != nil {
		return false
	}

	// check token expiry
	expTime, err := claims.GetExpirationTime()
	if err != nil || expTime.Before(time.Now()) {
		return false
	}

	return true
}

// plain text for unauthorized
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// welocme page for unauthorized
func AuthMiddlewareHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			welcomeapi.WelcomePage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
