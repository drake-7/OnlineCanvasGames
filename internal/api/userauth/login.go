package userauthapi

import (
	"database/sql"

	"github.com/1001bit/OnlineCanvasGames/internal/crypt"
	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
)

func login(userInput *WelcomeUserInput) (*usermodel.User, error) {
	user, hash, err := usermodel.GetUserAndHash(&userInput.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoUser
		}
		return nil, err
	}

	if !crypt.CheckHash(&userInput.Password, hash) {
		return nil, ErrNoUser
	}

	return user, nil
}
