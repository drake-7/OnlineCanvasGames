package userauthapi

import (
	"database/sql"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

func login(userInput *WelcomeUserInput) (auth.UserData, error) {
	userData := auth.UserData{Name: userInput.Username}

	// check user existance
	var hash string
	err := database.Statements["getHashAndId"].QueryRow(userInput.Username).Scan(&hash, &userData.ID)

	if err == sql.ErrNoRows || !auth.CheckHash(&userInput.Password, &hash) {
		return auth.UserData{}, ErrNoUser
	}

	if err != nil {
		return auth.UserData{}, err
	}

	return userData, err
}
