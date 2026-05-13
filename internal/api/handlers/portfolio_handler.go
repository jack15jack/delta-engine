package handlers

import (
	"net/http"
	"strconv"

	"github.com/jack15jack/delta-engine/internal/portfolio"

	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	service *portfolio.Service
}

func NewPortfolioHandler(s *portfolio.Service) *PortfolioHandler {
	return &PortfolioHandler{service: s}
}

func (h *PortfolioHandler) CreatePortfolio(c *gin.Context) {

	userID := c.Query("user_id")

	p, err := h.service.CreatePortfolio(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *PortfolioHandler) GetPortfolio(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	p, err := h.service.GetPortfolio(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *PortfolioHandler) Positions(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	pos, err := h.service.GetPositions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pos)
}

func (h *PortfolioHandler) Trade(c *gin.Context) {

	var req struct {
		PortfolioID int     `json:"portfolio_id"`
		Ticker      string  `json:"ticker"`
		Side        string  `json:"side"`
		Quantity    float64 `json:"quantity"`
		Price       float64 `json:"price"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ExecuteTrade(
		req.PortfolioID,
		req.Ticker,
		req.Side,
		req.Quantity,
		req.Price,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "executed"})
}
