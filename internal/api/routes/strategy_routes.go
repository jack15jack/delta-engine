package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jack15jack/delta-engine/internal/api/handlers"
	"github.com/jack15jack/delta-engine/internal/db"
	"github.com/jack15jack/delta-engine/internal/db/queries"
	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/strategy"
)

func RegisterStrategyRoutes(api *gin.RouterGroup) {

	provider := market.NewFinnhubProvider()
	marketService := market.NewService(provider)

	dbConn := db.NewPostgres()
	signalRepo := queries.NewSignalRepo(dbConn)

	sma := &strategy.SMACrossStrategy{
		Short: 5,
		Long:  20,
	}

	engine := strategy.NewEngine(sma)

	handler := handlers.NewStrategyHandler(marketService, engine, signalRepo)

	strategyGroup := api.Group("/strategy")

	strategyGroup.GET("/run/:symbol", handler.Run)
}
