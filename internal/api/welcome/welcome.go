package welcomeapi

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type WelcomeUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func WelcomePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")

	// decode request
	var inputData WelcomeUserInput
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Could not decode the request", http.StatusBadRequest)
		return
	}

	// disallow empty fields
	if inputData.Password == "" || inputData.Username == "" {
		http.Error(w, "Password or username is empty", http.StatusBadRequest)
		return
	}

	// disallow username with special characters
	if inputData.Username != regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(inputData.Username, "") {
		http.Error(w, "Username must not contain special characters", http.StatusBadRequest)
		return
	}

	// disallow short password
	if len(inputData.Password) < 8 {
		http.Error(w, "Password should be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// TODO: DATABASE
	http.Error(w, "Could not reach database", http.StatusInternalServerError)
}
