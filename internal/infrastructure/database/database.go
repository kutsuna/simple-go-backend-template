package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

func NewDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:./db/data/sample.db?cache=shared&mode=rwc")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	return db, nil
}
