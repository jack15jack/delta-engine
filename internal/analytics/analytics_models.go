package analytics

import "time"

type PortfolioPerformance struct {
	PortfolioID     uint    `json:"portfolio_id"`
	CashBalance     float64 `json:"cash_balance"`
	MarketValue     float64 `json:"market_value"`
	TotalValue      float64 `json:"total_value"`
	UnrealizedPnL   float64 `json:"unrealized_pnl"`
	TotalReturnPct  float64 `json:"total_return_pct"`
	BuyingPower     float64 `json:"buying_power"`
	PositionCount   int     `json:"position_count"`
	LongExposurePct float64 `json:"long_exposure_pct"`
	CashPct         float64 `json:"cash_pct"`
	LeverageRatio   float64 `json:"leverage_ratio"`
}

type ExposureMetrics struct {
	GrossExposure      float64 `json:"gross_exposure"`
	NetExposure        float64 `json:"net_exposure"`
	LongExposure       float64 `json:"long_exposure"`
	ShortExposure      float64 `json:"short_exposure"`
	LargestPosition    string  `json:"largest_position"`
	LargestPositionPct float64 `json:"largest_position_pct"`
	ConcentrationRisk  float64 `json:"concentration_risk"`
}

type EquitySnapshot struct {
	ID          uint
	PortfolioID uint
	Equity      float64
	Timestamp   time.Time
}

type PortfolioSnapshot struct {
	PortfolioID    uint
	CashBalance    float64
	MarketValue    float64
	TotalValue     float64
	UnrealizedPnL  float64
	PositionCount  int
	PositionValues map[string]float64
}

type RiskMetrics struct {
	PortfolioVolatility float64 `json:"portfolio_volatility"`
	ValueAtRisk95       float64 `json:"value_at_risk_95"`
	MaxPositionRisk     float64 `json:"max_position_risk"`
	CapitalAtRisk       float64 `json:"capital_at_risk"`
	RiskUtilization     float64 `json:"risk_utilization"`
}

type TradeMetrics struct {
	TotalTrades   int     `json:"total_trades"`
	WinningTrades int     `json:"winning_trades"`
	LosingTrades  int     `json:"losing_trades"`
	WinRate       float64 `json:"win_rate"`
	ProfitFactor  float64 `json:"profit_factor"`
	AveragePnL    float64 `json:"average_pnl"`
}

type PositionMetrics struct {
	TotalPositions        int     `json:"total_positions"`
	LargestPosition       string  `json:"largest_position"`
	AveragePositionSize   float64 `json:"average_position_size"`
	CashAllocationPct     float64 `json:"cash_allocation_pct"`
	InvestedAllocationPct float64 `json:"invested_allocation_pct"`
}
