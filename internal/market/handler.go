package market

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	service *Service
}

func NewMarketHandler(service *Service) *MarketHandler {
	return &MarketHandler{service: service}
}

func (h *MarketHandler) GetQuote(c *gin.Context) {
	symbol := c.Param("symbol")

	data, err := h.service.GetQuote(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
