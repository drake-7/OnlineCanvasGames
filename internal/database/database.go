package database

import (
	"database/sql"
	"fmt"

	"github.com/1001bit/OnlineCanvasGames/internal/env"
	_ "github.com/lib/pq"
)

var Database *sql.DB

type DatabaseConf struct {
	user string
	name string
	pass string
}

func InitDB() error {
	dbConf := DatabaseConf{
		env.GetEnv("DB_USER"),
		env.GetEnv("DB_NAME"),
		env.GetEnv("DB_PASS"),
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.user, dbConf.pass, dbConf.name)

	var err error
	Database, err = sql.Open("postgres", connStr)
	return err
}
