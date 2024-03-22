package usermodel

import (
	"github.com/1001bit/OnlineCanvasGames/internal/crypt"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

type User struct {
	ID   string
	Name string
}

func NameExists(username string) (bool, error) {
	// check user existance
	var exists bool

	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	return true, nil
}

func Insert(username, password string) (*User, error) {
	newUser := &User{Name: username}

	hash, err := crypt.GenerateHash(password)
	if err != nil {
		return nil, err
	}

	err = database.DB.QueryRow("INSERT INTO users (name, hash) VALUES ($1, $2) RETURNING id", username, hash).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func GetUserAndHash(username string) (*User, string, error) {
	user := &User{Name: username}

	// check user existance
	var hash string

	err := database.DB.QueryRow("SELECT id, hash FROM users WHERE name = $1", username).Scan(&user.ID, &hash)
	if err != nil {
		return nil, "", err
	}

	return user, hash, nil
}
