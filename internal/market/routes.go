package market

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMarketRoutes(api *gin.RouterGroup, dbConn *gorm.DB, service *Service) {
	handler := NewMarketHandler(service)

	marketGroup := api.Group("/market")

	marketGroup.GET("quote/:symbol", handler.GetQuote)
}
