package main

import (
	"fmt"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/router"
)

func main() {
	router := router.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on localhost%s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
