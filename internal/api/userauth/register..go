package userauthapi

import (
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

func register(userInput *WelcomeUserInput) (auth.UserData, error) {
	userData := auth.UserData{Name: userInput.Username}

	// check user existance
	var exists bool

	err := database.Statements["userExists"].QueryRow(userInput.Username).Scan(&exists)
	if err != nil {
		return auth.UserData{}, err
	}

	if exists {
		return auth.UserData{}, ErrUserExists
	}

	// create new user
	hash, err := auth.GenerateHash(&userInput.Password)
	if err != nil {
		return auth.UserData{}, err
	}

	err = database.Statements["register"].QueryRow(userInput.Username, hash).Scan(&userData.ID)
	if err != nil {
		return auth.UserData{}, err
	}

	return auth.UserData{}, nil
}
