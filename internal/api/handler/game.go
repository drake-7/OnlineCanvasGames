package handler

import (
	"fmt"
	"net/http"
)

func GamePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("game page %s", r.PathValue("id"))))
}
