package analytics

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/portfolio"
)

func RegisterAnalyticsRoutes(api *gin.RouterGroup, dbConn *gorm.DB, marketService *market.Service) {

	portfolioService := portfolio.NewService(dbConn)

	analyticsService := NewService(portfolioService, marketService)

	handler := NewHandler(analyticsService)

	a := api.Group("/analytics")

	a.GET("/portfolio/:id/performance", handler.GetPortfolioPerformance)
}
