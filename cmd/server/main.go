package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jack15jack/delta-engine/internal/analytics"
	"github.com/jack15jack/delta-engine/internal/backtest"
	"github.com/jack15jack/delta-engine/internal/config"
	"github.com/jack15jack/delta-engine/internal/db"
	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/middleware"
	"github.com/jack15jack/delta-engine/internal/portfolio"
	"github.com/jack15jack/delta-engine/internal/strategy"
)

func main() {

	cfg := config.LoadConfig()

	router := gin.New()

	// Core middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.Auth())

	// Register all route groups
	api := router.Group("/api")

	dbConn := db.NewSQLite()
	db.AutoMigrate(dbConn)

	provider := market.NewFinnhubProvider()
	marketService := market.NewService(provider)

	backtestService := backtest.NewService(marketService)
	backtestHandler := backtest.NewHandler(backtestService)

	market.RegisterMarketRoutes(api, dbConn, marketService)
	strategy.RegisterStrategyRoutes(api, dbConn)
	portfolio.RegisterPortfolioRoutes(api, dbConn)
	analytics.RegisterAnalyticsRoutes(api, dbConn, marketService)
	backtest.RegisterBacktestRoutes(api, backtestHandler)

	log.Printf("Delta Engine API running on :%s", cfg.Port)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
