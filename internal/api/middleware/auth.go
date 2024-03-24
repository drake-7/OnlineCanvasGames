package middleware

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/api/handler"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

// plain text for unauthorized
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.CheckCookieJWT(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// welcome page for unauthorized
func AuthHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.CheckCookieJWT(r) {
			handler.WelcomePage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
