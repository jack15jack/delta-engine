package backtest

import (
	"github.com/jack15jack/delta-engine/internal/models"
)

type Simulator struct {
	InitialCapital float64
	Cash           float64
	Position       *SimulatedPosition
	Trades         []models.Trade
	Snapshots      []EquitySnapshot
}

func NewSimulator(initialCapital float64) *Simulator {
	return &Simulator{
		InitialCapital: initialCapital,
		Cash:           initialCapital,
	}
}

func (s *Simulator) Buy(ticker string, price float64, qty float64) {

	cost := price * qty

	if cost > s.Cash {
		return
	}

	s.Cash -= cost

	s.Position = &SimulatedPosition{
		Ticker:   ticker,
		Quantity: qty,
		AvgCost:  price,
	}
}

func (s *Simulator) Sell(price float64) {

	if s.Position == nil {
		return
	}

	value := s.Position.Quantity * price

	s.Cash += value

	s.Position = nil
}

func (s *Simulator) Snapshot(candle models.Candle) {

	var marketValue float64

	if s.Position != nil {
		marketValue = s.Position.Quantity * candle.Close
	}

	s.Snapshots = append(
		s.Snapshots,
		EquitySnapshot{
			Time:        candle.Time,
			Cash:        s.Cash,
			MarketValue: marketValue,
			TotalValue:  s.Cash + marketValue,
		},
	)
}
