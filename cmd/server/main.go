package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1001bit/OnlineCanvasGames/internal/auth"
	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/env"
	"github.com/1001bit/OnlineCanvasGames/internal/router"
)

func init() {
	env.InitEnv()

	// init database
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	database.InitStatements()

	// init JWT
	auth.InitJWTSecret()
}

func main() {
	// close database on end
	defer database.Database.Close()

	// init http server
	router := router.NewRouter()

	port := 8080
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
