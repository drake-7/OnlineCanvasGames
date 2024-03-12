package userauthapi

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, userInput WelcomeUserInput) error {
	// TODO: DATABASE
	if false {
		return fmt.Errorf("%s doesn't exist", userInput.Username)
	}

	return nil
}
