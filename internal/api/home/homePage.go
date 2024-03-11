package homeapi

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page"))
}
