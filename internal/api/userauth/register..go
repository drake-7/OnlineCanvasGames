package userauthapi

import (
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

func register(userInput WelcomeUserInput) (string, error) {
	// check user existance
	var exists bool

	err := database.Statements["userExists"].QueryRow(userInput.Username).Scan(&exists)
	if err != nil {
		return "", err
	}

	if exists {
		return "", ErrUserExists
	}

	// create new user
	hash, err := auth.GenerateHash(userInput.Password)
	if err != nil {
		return "", err
	}

	var userID string
	err = database.Statements["register"].QueryRow(userInput.Username, hash).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}
