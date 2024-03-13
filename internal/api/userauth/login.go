package userauthapi

import (
	"database/sql"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

func login(userInput WelcomeUserInput) (string, error) {
	// check user existance
	var hash, id string
	err := database.Statements["getHashAndId"].QueryRow(userInput.Username).Scan(&hash, &id)

	if err == sql.ErrNoRows || !auth.CheckHash(userInput.Password, hash) {
		return "", ErrNoUser
	}

	if err != nil {
		return "", err
	}

	return id, err
}
