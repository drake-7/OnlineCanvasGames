package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/api/router"
)

func main() {
	router := router.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
