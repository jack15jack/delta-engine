package strategy

import (
	"github.com/jack15jack/delta-engine/internal/models"
	"github.com/jack15jack/delta-engine/internal/strategy/indicators"
)

type SMACrossStrategy struct {
	Short int
	Long  int
}

func (s *SMACrossStrategy) Name() string {
	return "sma_crossover"
}

func (s *SMACrossStrategy) Evaluate(
	candle models.Candle,
	history []models.Candle,
) *models.Signal {

	if len(history) < s.Long {
		return nil
	}

	shortPrices := []float64{}
	longPrices := []float64{}

	for _, c := range history {
		shortPrices = append(shortPrices, c.Close)
		longPrices = append(longPrices, c.Close)
	}

	shortSMA := indicators.SMA(shortPrices, s.Short)
	longSMA := indicators.SMA(longPrices, s.Long)

	if shortSMA == 0 || longSMA == 0 {
		return nil
	}

	// Signal logic
	if shortSMA > longSMA {
		return &models.Signal{
			Ticker:   candle.Ticker,
			Type:     models.BUY,
			Price:    candle.Close,
			Time:     candle.Time,
			Strategy: s.Name(),
			Strength: 0.7,
		}
	}

	if shortSMA < longSMA {
		return &models.Signal{
			Ticker:   candle.Ticker,
			Type:     models.SELL,
			Price:    candle.Close,
			Time:     candle.Time,
			Strategy: s.Name(),
			Strength: 0.7,
		}
	}

	return &models.Signal{
		Ticker:   candle.Ticker,
		Type:     models.HOLD,
		Price:    candle.Close,
		Time:     candle.Time,
		Strategy: s.Name(),
		Strength: 0.3,
	}
}
