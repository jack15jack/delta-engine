package db

import (
	"log"

	"github.com/jack15jack/delta-engine/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Candle{},
		&models.Portfolio{},
		&models.Position{},
		&models.Signal{},
		&models.Trade{},
	)

	if err != nil {
		log.Fatal("migration failed:", err)
	}
}
