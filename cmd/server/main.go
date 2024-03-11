package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/router"
)

func main() {
	router := router.NewRouter()

	port := 8080
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
