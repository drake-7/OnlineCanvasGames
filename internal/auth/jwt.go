package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/1001bit/OnlineCanvasGames/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

const JWTLifeTime = time.Hour * 24

var (
	ErrBadToken = errors.New("invalid token")
	secret      = []byte("")
)

func InitJWTSecret() {
	secret = []byte(env.GetEnvVal("JWT_SECRET"))
}

func CreateJWT(userID int, username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":   userID,
			"username": username,
			"exp":      time.Now().Add(JWTLifeTime).Unix(),
		},
	)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func JWTClaimsByCookie(r *http.Request) (jwt.MapClaims, error) {
	// get token from cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	// get token claims
	claims, err := jwtClaimsByString(cookie.Value)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func jwtByString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrBadToken
	}

	return token, nil
}

func jwtClaimsByString(tokenString string) (jwt.MapClaims, error) {
	token, err := jwtByString(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrBadToken
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil || expTime.Before(time.Now()) {
		return nil, ErrBadToken
	}

	return claims, nil
}
