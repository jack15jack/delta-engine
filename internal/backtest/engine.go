package backtest

import (
	"github.com/jack15jack/delta-engine/internal/models"
	"github.com/jack15jack/delta-engine/internal/strategy"
)

type Engine struct {
	strategy strategy.Strategy
}

func NewEngine(s strategy.Strategy) *Engine {
	return &Engine{
		strategy: s,
	}
}

func (e *Engine) Run(candles []models.Candle, initialCapital float64) (*Results, error) {

	sim := NewSimulator(initialCapital)

	history := []models.Candle{}

	for _, candle := range candles {

		history = append(history, candle)

		signal := e.strategy.Evaluate(candle, history)

		if signal != nil {

			switch signal.Type {

			case models.BUY:
				if sim.Position == nil {
					qty := sim.Cash / candle.Close
					sim.Buy(candle.Ticker, candle.Close, qty)
				}

			case models.SELL:
				if sim.Position != nil {
					sim.Sell(candle.Close)
				}
			}
		}

		sim.Snapshot(candle)
	}

	return sim.Results(), nil
}

func (s *Simulator) Results() *Results {

	finalValue := s.Cash

	if len(s.Snapshots) > 0 {
		finalValue = s.Snapshots[len(s.Snapshots)-1].TotalValue
	}

	initialCapital := s.InitialCapital

	totalReturnPct := ((finalValue - initialCapital) / initialCapital) * 100

	return &Results{
		InitialCapital: initialCapital,
		FinalValue:     finalValue,
		TotalReturnPct: totalReturnPct,
		EquityCurve:    s.Snapshots,
	}
}
