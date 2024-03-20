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

	stmt, err := database.DB.GetStatement("userExists")
	if err != nil {
		return false, err
	}
	err = stmt.QueryRow(username).Scan(&exists)
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

	stmt, err := database.DB.GetStatement("register")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(username, hash).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func GetUserAndHash(username string) (*User, *string, error) {
	user := &User{Name: username}

	// check user existance
	var hash *string
	stmt, err := database.DB.GetStatement("getUserAndHash")
	if err != nil {
		return nil, nil, err
	}
	err = stmt.QueryRow(username).Scan(&user.ID, &hash)
	if err != nil {
		return nil, nil, err
	}

	return user, hash, nil
}
