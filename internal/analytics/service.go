package analytics

import (
	"math"

	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/portfolio"
)

type Service struct {
	portfolio *portfolio.Service
	market    *market.Service
}

func NewService(
	portfolioService *portfolio.Service,
	marketService *market.Service,
) *Service {
	return &Service{
		portfolio: portfolioService,
		market:    marketService,
	}
}

func (s *Service) BuildSnapshot(portfolioID int) (*PortfolioSnapshot, error) {

	p, err := s.portfolio.GetPortfolio(portfolioID)
	if err != nil {
		return nil, err
	}

	positions, err := s.portfolio.GetPositions(portfolioID)
	if err != nil {
		return nil, err
	}

	snapshot := &PortfolioSnapshot{
		PortfolioID:    p.ID,
		CashBalance:    p.CashBalance,
		PositionCount:  len(positions),
		PositionValues: make(map[string]float64),
	}

	for _, pos := range positions {

		quote, err := s.market.GetQuote(pos.Ticker)
		if err != nil {
			continue
		}

		currentPrice := quote.Close

		positionValue := pos.Quantity * currentPrice

		unrealizedPnL :=
			(currentPrice - pos.AvgCost) * pos.Quantity

		snapshot.MarketValue += positionValue
		snapshot.UnrealizedPnL += unrealizedPnL

		snapshot.PositionValues[pos.Ticker] = positionValue
	}

	snapshot.TotalValue =
		snapshot.CashBalance + snapshot.MarketValue

	return snapshot, nil
}

func (s *Service) GetPortfolioPerformance(portfolioID int) (*PortfolioPerformance, error) {

	snapshot, err := s.BuildSnapshot(portfolioID)
	if err != nil {
		return nil, err
	}

	initialCapital := 100000.0

	totalReturnPct :=
		((snapshot.TotalValue - initialCapital) / initialCapital) * 100

	var cashPct float64
	var longExposurePct float64
	var leverageRatio float64

	if snapshot.TotalValue > 0 {

		cashPct =
			(snapshot.CashBalance / snapshot.TotalValue) * 100

		longExposurePct =
			(snapshot.MarketValue / snapshot.TotalValue) * 100

		leverageRatio =
			snapshot.MarketValue / snapshot.TotalValue
	}

	return &PortfolioPerformance{
		PortfolioID:     snapshot.PortfolioID,
		CashBalance:     snapshot.CashBalance,
		MarketValue:     snapshot.MarketValue,
		TotalValue:      snapshot.TotalValue,
		UnrealizedPnL:   snapshot.UnrealizedPnL,
		TotalReturnPct:  totalReturnPct,
		BuyingPower:     snapshot.CashBalance,
		PositionCount:   snapshot.PositionCount,
		LongExposurePct: longExposurePct,
		CashPct:         cashPct,
		LeverageRatio:   leverageRatio,
	}, nil
}

func (s *Service) GetExposure(portfolioID int) (*ExposureMetrics, error) {

	snapshot, err := s.BuildSnapshot(portfolioID)
	if err != nil {
		return nil, err
	}

	var largestTicker string
	var largestValue float64

	for ticker, value := range snapshot.PositionValues {

		if value > largestValue {
			largestValue = value
			largestTicker = ticker
		}
	}

	var largestPct float64

	if snapshot.MarketValue > 0 {
		largestPct =
			(largestValue / snapshot.MarketValue) * 100
	}

	return &ExposureMetrics{
		GrossExposure:      snapshot.MarketValue,
		NetExposure:        snapshot.MarketValue,
		LongExposure:       snapshot.MarketValue,
		ShortExposure:      0,
		LargestPosition:    largestTicker,
		LargestPositionPct: largestPct,
		ConcentrationRisk:  largestPct,
	}, nil
}

func (s *Service) GetRiskMetrics(portfolioID int) (*RiskMetrics, error) {

	snapshot, err := s.BuildSnapshot(portfolioID)
	if err != nil {
		return nil, err
	}

	var maxPosition float64

	for _, value := range snapshot.PositionValues {
		if value > maxPosition {
			maxPosition = value
		}
	}

	var maxPositionRisk float64
	var riskUtilization float64

	if snapshot.TotalValue > 0 {

		maxPositionRisk = (maxPosition / snapshot.TotalValue) * 100

		riskUtilization = (snapshot.MarketValue / snapshot.TotalValue) * 100
	}

	// Placeholder approximation until
	// historical returns exist

	portfolioVolatility := riskUtilization * 0.15

	valueAtRisk95 := snapshot.MarketValue * 0.02

	return &RiskMetrics{
		PortfolioVolatility: portfolioVolatility,
		ValueAtRisk95:       valueAtRisk95,
		MaxPositionRisk:     maxPositionRisk,
		CapitalAtRisk:       snapshot.MarketValue,
		RiskUtilization:     riskUtilization,
	}, nil
}

func (s *Service) GetTradeMetrics(portfolioID int) (*TradeMetrics, error) {

	trades, err := s.portfolio.GetTrades(portfolioID)

	if err != nil {
		return nil, err
	}

	totalTrades := len(trades)

	if totalTrades == 0 {
		return &TradeMetrics{}, nil
	}

	var winningTrades int
	var losingTrades int

	var grossProfit float64
	var grossLoss float64

	for _, t := range trades {

		// Placeholder approximation

		if t.Side == "SELL" {

			pnl := t.Quantity * t.Price * 0.01

			if pnl > 0 {
				winningTrades++
				grossProfit += pnl
			} else {
				losingTrades++
				grossLoss += pnl
			}
		}
	}

	var winRate float64

	winRate = (float64(winningTrades) / float64(totalTrades)) * 100

	var profitFactor float64

	if grossLoss != 0 {
		profitFactor = grossProfit / math.Abs(grossLoss)
	}

	averagePnL := (grossProfit + grossLoss) / float64(totalTrades)

	return &TradeMetrics{
		TotalTrades:   totalTrades,
		WinningTrades: winningTrades,
		LosingTrades:  losingTrades,
		WinRate:       winRate,
		ProfitFactor:  profitFactor,
		AveragePnL:    averagePnL,
	}, nil
}

func (s *Service) GetPositionMetrics(portfolioID int) (*PositionMetrics, error) {

	snapshot, err := s.BuildSnapshot(portfolioID)
	if err != nil {
		return nil, err
	}

	var largestTicker string
	var largestValue float64

	for ticker, value := range snapshot.PositionValues {

		if value > largestValue {
			largestValue = value
			largestTicker = ticker
		}
	}

	var avgPosition float64

	if snapshot.PositionCount > 0 {
		avgPosition =
			snapshot.MarketValue /
				float64(snapshot.PositionCount)
	}

	var cashPct float64
	var investedPct float64

	if snapshot.TotalValue > 0 {

		cashPct =
			(snapshot.CashBalance / snapshot.TotalValue) * 100

		investedPct =
			(snapshot.MarketValue / snapshot.TotalValue) * 100
	}

	return &PositionMetrics{
		TotalPositions: snapshot.PositionCount,

		LargestPosition: largestTicker,

		AveragePositionSize: avgPosition,

		CashAllocationPct: cashPct,

		InvestedAllocationPct: investedPct,
	}, nil
}
