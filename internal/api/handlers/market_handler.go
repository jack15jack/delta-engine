package handlers

import (
	"net/http"

	"github.com/jack15jack/delta-engine/internal/market"

	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	service *market.Service
}

func NewMarketHandler(service *market.Service) *MarketHandler {
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
