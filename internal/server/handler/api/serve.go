package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ServeJSON(data any, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("err on response:", err)
		return
	}
	w.Write(b)
}

func ServeJSONMessage(message string, status int, w http.ResponseWriter) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	ServeJSON(data, status, w)
}
