package usermodel

import (
	"database/sql"
	"errors"

	"github.com/1001bit/OnlineCanvasGames/internal/crypt"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
)

type User struct {
	ID   int
	Name string
	Date string
}

var (
	ErrNoUserExists = errors.New("user with such name doesn't exist")
	ErrNoSuchUser   = errors.New("incorrect username or password")
	ErrUserExists   = errors.New("user with such name already exists")
)

func GetByID(userID int) (*User, error) {
	user := &User{ID: userID}

	err := database.DB.QueryRow("SELECT name, date FROM users WHERE id = $1", userID).Scan(&user.Name, &user.Date)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetByNameAndPassword(username, password string) (*User, error) {
	user := &User{Name: username}
	var hash string

	err := database.DB.QueryRow("SELECT id, date, hash FROM users WHERE name = $1", username).Scan(&user.ID, &user.Date, &hash)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNoSuchUser
		default:
			return nil, err
		}
	}

	if !crypt.CheckHash(password, hash) {
		return nil, ErrNoSuchUser
	}

	return user, nil
}

func Insert(username, password string) (*User, error) {
	// check existance
	var exists bool

	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)", username).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserExists
	}

	// create new user
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
