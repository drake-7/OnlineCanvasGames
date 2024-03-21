package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/1001bit/OnlineCanvasGames/internal/env"
	_ "github.com/lib/pq"
)

type DBConf struct {
	user string
	name string
	pass string
}

var (
	ErrNoStmt = errors.New("no statement found")
	DB        *sql.DB
)

func Start() error {
	// init database
	dbConf := DBConf{
		env.GetEnv("DB_USER"),
		env.GetEnv("DB_NAME"),
		env.GetEnv("DB_PASS"),
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.user, dbConf.pass, dbConf.name)

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
