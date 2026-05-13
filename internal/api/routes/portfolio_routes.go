package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jack15jack/delta-engine/internal/api/handlers"
	"github.com/jack15jack/delta-engine/internal/db"
	"github.com/jack15jack/delta-engine/internal/portfolio"
)

func RegisterPortfolioRoutes(api *gin.RouterGroup) {

	dbConn := db.NewPostgres()

	portfolioService := portfolio.NewService(dbConn)
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)

	p := api.Group("/portfolio")

	p.POST("/create", portfolioHandler.CreatePortfolio)
	p.GET("/:id", portfolioHandler.GetPortfolio)
	p.GET("/:id/positions", portfolioHandler.Positions)
	p.POST("/trade", portfolioHandler.Trade)
}
