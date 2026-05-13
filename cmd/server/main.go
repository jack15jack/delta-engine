package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jack15jack/delta-engine/internal/api/middleware"
	"github.com/jack15jack/delta-engine/internal/api/routes"
	"github.com/jack15jack/delta-engine/internal/config"
)

func main() {

	cfg := config.LoadConfig()

	router := gin.New()

	// Core middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())

	// Register all route groups
	routes.RegisterRoutes(router)

	log.Printf("Delta Engine API running on :%s", cfg.Port)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
