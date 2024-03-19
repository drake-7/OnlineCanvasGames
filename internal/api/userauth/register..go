package userauthapi

import (
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/model"
)

func register(userInput *WelcomeUserInput) (*model.User, error) {
	userData := &model.User{Name: userInput.Username}

	// check user existance
	var exists bool

	stmt, err := database.DB.GetStatement("userExist")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(userInput.Username).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUserExists
	}

	// create new user
	hash, err := auth.GenerateHash(&userInput.Password)
	if err != nil {
		return nil, err
	}

	stmt, err = database.DB.GetStatement("register")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(userInput.Username, hash).Scan(&userData.ID)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
