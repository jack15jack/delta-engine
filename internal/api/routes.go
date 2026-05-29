package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jack15jack/delta-engine/internal/analytics"
	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/portfolio"
	"github.com/jack15jack/delta-engine/internal/strategy"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, dbConn *gorm.DB) {

	api := router.Group("/api")

	provider := market.NewFinnhubProvider()
	marketService := market.NewService(provider)

	// health check first
	api.GET("/health", HealthCheck)

	market.RegisterMarketRoutes(api, dbConn, marketService)
	strategy.RegisterStrategyRoutes(api, dbConn)
	portfolio.RegisterPortfolioRoutes(api, dbConn)
	analytics.RegisterAnalyticsRoutes(api, dbConn, marketService)
}
