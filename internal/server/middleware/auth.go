package middleware

import (
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/api"
	"github.com/1001bit/OnlineCanvasGames/internal/server/handler/page"
)

type MessageJSON struct {
	message string
}

// plain text for unauthorized
func AuthJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.JWTClaimsByCookie(r)
		if err != nil {
			message := MessageJSON{
				message: "unauthorized",
			}
			api.ServeJSON(message, http.StatusUnauthorized, w)
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
			page.HandleAuth(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
