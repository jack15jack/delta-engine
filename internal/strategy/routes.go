package strategy

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jack15jack/delta-engine/internal/db/repos"
	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/strategy/strategies"
)

func RegisterStrategyRoutes(api *gin.RouterGroup, dbConn *gorm.DB) {

	provider := market.NewFinnhubProvider()
	marketService := market.NewService(provider)

	signalRepo := repos.NewSignalRepo(dbConn)

	sma := &strategies.SMACrossStrategy{
		Short: 5,
		Long:  20,
	}

	engine := NewEngine(sma)

	handler := NewStrategyHandler(marketService, engine, signalRepo)

	strategyGroup := api.Group("/strategy")

	strategyGroup.GET("/run/:symbol", handler.Run)
}
