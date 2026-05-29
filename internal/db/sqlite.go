package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() *gorm.DB {

	// Build path: internal/db/data/delta.db
	dbDir := filepath.Join("internal", "db", "data")

	// Ensure directory exists
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatal("failed to create db directory:", err)
	}

	dbPath := filepath.Join(dbDir, "delta.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to sqlite:", err)
	}

	return db
}
