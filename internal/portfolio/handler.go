package portfolio

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	service *Service
}

func NewPortfolioHandler(s *Service) *PortfolioHandler {
	return &PortfolioHandler{service: s}
}

func (h *PortfolioHandler) CreatePortfolio(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	p, err := h.service.CreatePortfolio(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, p)
}

func (h *PortfolioHandler) GetPortfolio(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "invalid portfolio id"})
		return
	}

	p, err := h.service.GetPortfolio(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "portfolio not found"})
		return
	}

	c.JSON(200, p)
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
