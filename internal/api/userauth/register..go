package userauthapi

import (
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/model"
)

func register(userInput *WelcomeUserInput) (model.User, error) {
	userData := model.User{Name: userInput.Username}

	// check user existance
	var exists bool

	err := database.Statements["userExists"].QueryRow(userInput.Username).Scan(&exists)
	if err != nil {
		return model.User{}, err
	}

	if exists {
		return model.User{}, ErrUserExists
	}

	// create new user
	hash, err := auth.GenerateHash(&userInput.Password)
	if err != nil {
		return model.User{}, err
	}

	err = database.Statements["register"].QueryRow(userInput.Username, hash).Scan(&userData.ID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{}, nil
}
