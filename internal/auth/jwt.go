package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret      = []byte("secret")
	JWTLifeTime = time.Hour * 24
	ErrBadToken = fmt.Errorf("invalid token")
)

func CreateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(JWTLifeTime).Unix(),
		},
	)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GetJWTClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrBadToken
}
