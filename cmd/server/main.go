package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/api/router"
	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/env"
)

func init() {
	env.InitEnv()
	auth.InitJWTSecret()
}

func main() {
	// start database
	err := database.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	// start http server
	router := router.NewRouter()

	port := 8080
	addr := fmt.Sprintf("localhost:%d", port)

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
