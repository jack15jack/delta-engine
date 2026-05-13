package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api")

	// health check first
	api.GET("/health", HealthCheck)

	RegisterMarketRoutes(api)
	RegisterStrategyRoutes(api)
	RegisterPortfolioRoutes(api)
	RegisterAnalyticsRoutes(api)
}
