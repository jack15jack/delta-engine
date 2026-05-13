package models

import "time"

type Trade struct {
	ID          int       `json:"id"`
	PortfolioID int       `json:"portfolio_id"`
	Ticker      string    `json:"ticker"`
	Side        string    `json:"side"`
	Quantity    float64   `json:"quantity"`
	Price       float64   `json:"price"`
	ExecutedAt  time.Time `json:"executed_at"`
}
