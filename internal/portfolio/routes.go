package portfolio

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPortfolioRoutes(api *gin.RouterGroup, dbConn *gorm.DB) {

	portfolioService := NewService(dbConn)
	portfolioHandler := NewPortfolioHandler(portfolioService)

	p := api.Group("/portfolio")

	p.POST("", portfolioHandler.CreatePortfolio)
	p.GET("/:id", portfolioHandler.GetPortfolio)
	p.GET("/:id/positions", portfolioHandler.Positions)
	p.POST("/trade", portfolioHandler.Trade)
}
