package database

import (
	_ "modernc.org/sqlite"

	"database/sql"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS short_urls (
			id TEXT NOT NULL,
			redirect_url TEXT NOT NULL,
			delete_id TEXT NOT NULL,
			expired_at INTEGER NOT NULL
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}
