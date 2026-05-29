package models

type Position struct {
	ID          int     `json:"id"`
	PortfolioID int     `json:"portfolio_id"`
	Ticker      string  `json:"ticker"`
	Quantity    float64 `json:"quantity"`
	AvgCost     float64 `json:"avg_cost"`
}
