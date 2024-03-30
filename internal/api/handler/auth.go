package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	usermodel "github.com/1001bit/OnlineCanvasGames/internal/model/user"
	"github.com/1001bit/OnlineCanvasGames/internal/tmplloader"
)

type AuthUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	// only unauthorized may see this page
	_, err := auth.JWTClaimsByCookie(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmplloader.ExecuteTemplate(w, r, "auth.html", nil)
}

func AuthPost(w http.ResponseWriter, r *http.Request) {
	// decode request
	var userInput AuthUserInput

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		ServerError(w, r)
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
	var user *usermodel.User
	switch userInput.Type {
	case "login":
		user, err = usermodel.GetByNameAndPassword(userInput.Username, userInput.Password)
	case "register":
		user, err = usermodel.Insert(userInput.Username, userInput.Password)
	}

	if err != nil {
		switch err {
		case usermodel.ErrNoSuchUser:
			http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		case usermodel.ErrUserExists:
			http.Error(w, fmt.Sprintf("%s already exists", userInput.Username), http.StatusUnauthorized)
		default:
			ServerError(w, r)
			log.Println("login/register err:", err)
		}
		return
	}

	// set token cookie
	token, err := auth.CreateJWT(user.ID, user.Name)
	if err != nil {
		ServerError(w, r)
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
