package handler

import (
	"net/http"
)

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func ServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong", http.StatusUnauthorized)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusUnauthorized)
}
