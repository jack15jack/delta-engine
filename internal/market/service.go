package market

import "github.com/jack15jack/delta-engine/internal/models"

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
		return nil, err
	}

	s.history[symbol] = append(s.history[symbol], *candle)

	return candle, nil
}

func (s *Service) GetHistory(symbol string) []models.Candle {
	return s.history[symbol]
}
