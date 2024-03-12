package userauthapi

import (
	"fmt"
	"net/http"
)

func register(w http.ResponseWriter, userInput WelcomeUserInput) error {
	// TODO: DATABASE
	if false {
		return fmt.Errorf("%s already exists", userInput.Username)
	}

	return nil
}
