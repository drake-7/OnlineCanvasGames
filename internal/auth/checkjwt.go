package auth

import (
	"net/http"
	"time"
)

func CheckCookieJWT(r *http.Request) bool {
	// get token from cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return false
	}

	// get token claims
	claims, err := GetJWTClaims(cookie.Value)
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
