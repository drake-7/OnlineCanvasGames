package page

import "net/http"

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	serveTemplate("notFound.html", nil, w, r)
}

func HandleServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	serveTemplate("serverError.html", nil, w, r)
}
