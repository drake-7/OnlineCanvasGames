package userauthapi

import (
	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
)

func register(userInput *WelcomeUserInput) (*usermodel.User, error) {
	// check user existance
	exists, err := usermodel.NameExists(&userInput.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserExists
	}

	// create new user
	userData, err := usermodel.Insert(&userInput.Username, &userInput.Password)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
