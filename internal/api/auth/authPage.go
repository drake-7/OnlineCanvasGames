package authapi

import "net/http"

func AuthPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth page"))
}
