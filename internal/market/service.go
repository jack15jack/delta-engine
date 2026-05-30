package market

import (
	"fmt"
	"time"

	"github.com/jack15jack/delta-engine/internal/models"
)

type Service struct {
	provider Provider
	history  map[string][]models.Candle
}

func NewService(p Provider) *Service {
	return &Service{
		provider: p,
		history:  make(map[string][]models.Candle),
	}
}

func (s *Service) GetQuote(symbol string) (*models.Candle, error) {

	candle, err := s.provider.GetQuote(symbol)
	if err != nil {
		return nil, fmt.Errorf("market.GetQuote(%s): %w", symbol, err)
	}

	s.history[symbol] = append(s.history[symbol], *candle)

	return candle, nil
}

func (s *Service) GetCachedHistory(symbol string) []models.Candle {
	return s.history[symbol]
}

func (s *Service) GetHistoricalData(symbol string, start, end time.Time) ([]models.Candle, error) {
	candles, err := s.provider.GetHistoricalData(symbol, start, end)
	if err != nil {
		return nil, fmt.Errorf("market.GetHistoricalData(%s): %w", symbol, err)
	}
	return candles, nil
}
