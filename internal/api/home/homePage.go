package homeapi

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page"))
}
