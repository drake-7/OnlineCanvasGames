package usermodel

import (
	"database/sql"
	"errors"

	"github.com/1001bit/OnlineCanvasGames/internal/crypt"
)

var (
	ErrUserWrong  = errors.New("incorrect username or password")
	ErrUserExists = errors.New("user with such name already exists")
)

func Login(username, password string) (*User, error) {
	user, hash, err := GetUserAndHash(username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserWrong
		}
		return nil, err
	}

	if !crypt.CheckHash(password, hash) {
		return nil, ErrUserWrong
	}

	return user, nil
}

func Register(username, password string) (*User, error) {
	// check user existance
	err := NameExists(username)
	switch err {
	case nil:
		return nil, ErrUserExists
	case ErrNoUserExists:
		break
	default:
		return nil, err
	}

	// create new user
	userData, err := Insert(username, password)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
