package storage

import "net/http"

const staticStoragePath = "./web/static"

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(staticStoragePath))
	fileServer.ServeHTTP(w, r)
}
