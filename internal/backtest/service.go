package backtest

import (
	"fmt"
	"time"

	"github.com/jack15jack/delta-engine/internal/market"
	"github.com/jack15jack/delta-engine/internal/strategy"
	"github.com/jack15jack/delta-engine/internal/strategy/strategies"
)

type Service struct {
	market *market.Service
}

func NewService(marketService *market.Service) *Service {
	return &Service{
		market: marketService,
	}
}

func (s *Service) Run(symbol string, start time.Time, end time.Time, strat strategy.Strategy, initialCapital float64) (*Results, error) {

	candles, err := s.market.GetHistoricalData(symbol, start, end)

	if err != nil {
		return nil, err
	}

	engine := NewEngine(strat)

	return engine.Run(candles, initialCapital)
}

func BuildStrategy(name string) (strategy.Strategy, error) {

	switch name {

	case "sma":
		return &strategies.SMACrossStrategy{
			Short: 20,
			Long:  50,
		}, nil

	default:
		return nil, fmt.Errorf("unknown strategy: %s", name)
	}
}
