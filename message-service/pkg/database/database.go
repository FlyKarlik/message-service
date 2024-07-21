package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ConnectToPostgres(databaseURL string) (*sqlx.DB, error) {
	conStr := fmt.Sprint(databaseURL)
	db, err := sqlx.Open("postgres", conStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
