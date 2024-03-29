package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/1001bit/OnlineCanvasGames/internal/env"
	_ "github.com/lib/pq"
)

var ErrNoStmt = errors.New("no statement found")

type DBConf struct {
	user string
	name string
	pass string
}

var DB *sql.DB

func Start() error {
	// init database
	dbConf := DBConf{
		env.GetEnvVal("DB_USER"),
		env.GetEnvVal("DB_NAME"),
		env.GetEnvVal("DB_PASS"),
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

func FormatPostgresDate(dateStr string) (string, error) {
	t, err := time.Parse("2006-01-02T15:04:05Z", dateStr)
	if err != nil {
		return "", err
	}

	formatted := t.Format("2 January 2006")
	return formatted, nil
}
