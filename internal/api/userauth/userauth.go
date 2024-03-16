package userauthapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
)

var (
	ErrNoUser     = fmt.Errorf("incorrect username or password")
	ErrUserExists = fmt.Errorf("user with such name already exists")
)

type WelcomeUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func UserAuthPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")

	// decode request
	var userInput WelcomeUserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Could not decode the request", http.StatusBadRequest)
		return
	}

	// disallow empty fields
	if userInput.Password == "" || userInput.Username == "" {
		http.Error(w, "Password or username is empty", http.StatusBadRequest)
		return
	}

	// disallow username with special characters
	if userInput.Username != regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(userInput.Username, "") {
		http.Error(w, "Username must not contain special characters", http.StatusBadRequest)
		return
	}

	// disallow short password
	if len(userInput.Password) < 8 {
		http.Error(w, "Password should be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Login / register
	var userData auth.UserData
	if userInput.Type == "login" {
		userData, err = login(&userInput)
	} else {
		userData, err = register(&userInput)
	}

	if err != nil {
		switch err {
		case ErrNoUser:
			http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		case ErrUserExists:
			http.Error(w, fmt.Sprintf("%s already exists", userInput.Username), http.StatusUnauthorized)
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
			log.Println("login/register err:", err)
		}
		return
	}

	// set token cookie
	token, err := auth.CreateJWT(&userData)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		log.Println("jwt creation err:", err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		MaxAge:   int(auth.JWTLifeTime.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	w.Write([]byte("Success!"))
}
