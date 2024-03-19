package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/1001bit/OnlineCanvasGames/internal/env"
	_ "github.com/lib/pq"
)

var (
	ErrNoStmt = errors.New("no statement found")
	DB        *Database
)

type DBConf struct {
	user string
	name string
	pass string
}

type Database struct {
	db         *sql.DB
	statements map[string]*sql.Stmt
}

func (database *Database) Start() error {
	database.statements = make(map[string]*sql.Stmt)

	// init database
	dbConf := DBConf{
		env.GetEnv("DB_USER"),
		env.GetEnv("DB_NAME"),
		env.GetEnv("DB_PASS"),
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.user, dbConf.pass, dbConf.name)

	var err error

	database.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = database.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) Close() error {
	return database.db.Close()
}

func (database *Database) prepareStatement(name, statement string) {
	var err error
	database.statements[name], err = database.db.Prepare(statement)
	if err != nil {
		log.Println(name, "statement error:", err)
	}
}

func (database *Database) GetStatement(name string) (*sql.Stmt, error) {
	stmt, ok := database.statements[name]
	if !ok {
		return nil, ErrNoStmt
	}
	return stmt, nil
}
