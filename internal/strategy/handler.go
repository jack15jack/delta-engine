package strategy

import (
	"net/http"

	"github.com/jack15jack/delta-engine/internal/db/repos"
	"github.com/jack15jack/delta-engine/internal/market"

	"github.com/gin-gonic/gin"
)

type StrategyHandler struct {
	marketService *market.Service
	stratEngine   *Engine
	signalRepo    *repos.SignalRepo
}

func NewStrategyHandler(
	ms *market.Service,
	se *Engine,
	repo *repos.SignalRepo,
) *StrategyHandler {
	return &StrategyHandler{
		marketService: ms,
		stratEngine:   se,
		signalRepo:    repo,
	}
}

func (h *StrategyHandler) Run(c *gin.Context) {

	symbol := c.Param("symbol")

	candle, err := h.marketService.GetQuote(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	history := h.marketService.GetHistory(symbol)

	signals := h.stratEngine.Run(*candle, history)

	// store signals in db
	for _, sig := range signals {
		_ = h.signalRepo.Insert(*sig)
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol":  symbol,
		"signals": signals,
	})
}
