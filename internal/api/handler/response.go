package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(response any, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("err on response:", err)
	}
	w.Write(b)
}
