package backtest

import "github.com/gin-gonic/gin"

func RegisterBacktestRoutes(api *gin.RouterGroup, handler *Handler) {
	bt := api.Group("/backtest")
	bt.POST("/run", handler.Run)
}
