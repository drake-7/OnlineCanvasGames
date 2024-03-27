package handler

import "net/http"

const staticStoragePath = "./web/static"

func StaticStorage(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(staticStoragePath))
	fileServer.ServeHTTP(w, r)
}
