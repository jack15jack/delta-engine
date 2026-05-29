package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jack15jack/delta-engine/internal/api"
	"github.com/jack15jack/delta-engine/internal/config"
	"github.com/jack15jack/delta-engine/internal/db"
	"github.com/jack15jack/delta-engine/internal/middleware"
)

func main() {

	cfg := config.LoadConfig()

	router := gin.New()

	// Core middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.Auth())

	dbConn := db.NewSQLite()
	db.AutoMigrate(dbConn)

	// Register all route groups
	api.RegisterRoutes(router, dbConn)

	log.Printf("Delta Engine API running on :%s", cfg.Port)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
