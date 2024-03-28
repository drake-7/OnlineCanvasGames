package middleware

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/api/handler"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

// plain text for unauthorized
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.JWTClaimsByCookie(r)
		if err != nil {
			handler.Unauthorized(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// auth page for unauthorized
func AuthHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.JWTClaimsByCookie(r)
		if err != nil {
			handler.AuthPage(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
