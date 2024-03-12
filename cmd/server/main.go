package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/1001bit/OnlineCanvasGames/internal/database"
	"github.com/1001bit/OnlineCanvasGames/internal/env"
	"github.com/1001bit/OnlineCanvasGames/internal/router"
)

func init() {
	env.InitEnv()

	// init database
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	// close database on end
	defer database.Database.Close()

	// TEST CASE
	var id int
	err := database.Database.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", time.Now().Unix()).Scan(&id)
	if err != nil {
		panic(err)
	}
	log.Println(id)

	// init http server
	router := router.NewRouter()

	port := 8080
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
