package backtest

import (
	"time"
)

type EquitySnapshot struct {
	Time        time.Time `json:"time"`
	Cash        float64   `json:"cash"`
	MarketValue float64   `json:"market_value"`
	TotalValue  float64   `json:"total_value"`
}

type SimulatedPosition struct {
	Ticker   string
	Quantity float64
	AvgCost  float64
}

type Results struct {
	InitialCapital float64          `json:"initial_capital"`
	FinalValue     float64          `json:"final_value"`
	TotalReturnPct float64          `json:"total_return_pct"`
	TotalTrades    int              `json:"total_trades"`
	WinningTrades  int              `json:"winning_trades"`
	LosingTrades   int              `json:"losing_trades"`
	EquityCurve    []EquitySnapshot `json:"equity_curve"`
}

type Request struct {
	Symbol         string    `json:"symbol"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	InitialCapital float64   `json:"initial_capital"`
	Strategy       string    `json:"strategy"`
}
