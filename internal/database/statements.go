package database

import (
	"database/sql"
	"log"
)

var Statements = make(map[string]*sql.Stmt)

func prepareStatement(name, statement string) {
	var err error
	Statements[name], err = Database.Prepare(statement)
	log.Println(name, "statement error:", err)
}

func InitStatements() {
	// AUTH
	// register
	prepareStatement("register", "INSERT INTO users (name, hash) VALUES ($1, $2) RETURNING id")
	// user existance
	prepareStatement("userExists", "SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)")
	// hash and user id
	prepareStatement("getHashAndId", "SELECT hash, id FROM users WHERE name = $1")

	// GAMES
	// full list
	prepareStatement("getGames", "SELECT id, title FROM games")
}
