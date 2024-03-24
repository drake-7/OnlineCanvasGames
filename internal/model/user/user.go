package usermodel

import (
	"errors"

	"github.com/1001bit/OnlineCanvasGames/internal/crypt"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

type User struct {
	ID   string
	Name string
}

var (
	ErrNoUserExists = errors.New("user with such name doesn't exist")
)

func NameExists(username string) error {
	var exists bool

	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)", username).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return ErrNoUserExists
	}

	return nil
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
	var hash string

	err := database.DB.QueryRow("SELECT id, hash FROM users WHERE name = $1", username).Scan(&user.ID, &hash)
	if err != nil {
		return nil, "", err
	}

	return user, hash, nil
}
