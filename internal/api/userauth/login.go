package userauthapi

import (
	"database/sql"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/model"
)

func login(userInput *WelcomeUserInput) (*model.User, error) {
	userData := &model.User{Name: userInput.Username}

	// check user existance
	var hash string
	stmt, err := database.DB.GetStatement("getHashAndId")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(userInput.Username).Scan(&hash, &userData.ID)

	if err == sql.ErrNoRows || !auth.CheckHash(&userInput.Password, &hash) {
		return nil, ErrNoUser
	}

	if err != nil {
		return nil, err
	}

	return userData, err
}
