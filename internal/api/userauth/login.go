package userauthapi

import (
	"database/sql"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/model"
)

func login(userInput *WelcomeUserInput) (model.User, error) {
	userData := model.User{Name: userInput.Username}

	// check user existance
	var hash string
	err := database.Statements["getHashAndId"].QueryRow(userInput.Username).Scan(&hash, &userData.ID)

	if err == sql.ErrNoRows || !auth.CheckHash(&userInput.Password, &hash) {
		return model.User{}, ErrNoUser
	}

	if err != nil {
		return model.User{}, err
	}

	return userData, err
}
