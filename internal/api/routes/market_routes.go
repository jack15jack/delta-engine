package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jack15jack/delta-engine/internal/api/handlers"
	"github.com/jack15jack/delta-engine/internal/market"
)

func RegisterMarketRoutes(api *gin.RouterGroup) {
	provider := market.NewFinnhubProvider()
	service := market.NewService(provider)
	handler := handlers.NewMarketHandler(service)

	marketGroup := api.Group("/market")

	marketGroup.GET("quote/:symbol", handler.GetQuote)
}
