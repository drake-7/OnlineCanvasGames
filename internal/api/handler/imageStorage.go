package handler

import (
	"net/http"
	"os"
)

const imageStoragePath = "./web/image"

var imageFormats = [...]string{"png", "jpg", "jpeg"}

func ImageStorage(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(imageStoragePath))

	// check if file.png, file.jpg, etc. exists and use it
	for _, format := range imageFormats {
		possiblePath := imageStoragePath + r.URL.Path + "." + format

		if _, err := os.Stat(possiblePath); err == nil {
			r.URL.Path += "." + format
			break
		}
	}

	fileServer.ServeHTTP(w, r)
}
