package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgres() *sql.DB {

	dsn := os.Getenv("POSTGRES_DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	return db
}
