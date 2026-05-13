package strategy

import "github.com/jack15jack/delta-engine/internal/models"

type Engine struct {
	strategies []Strategy
}

func NewEngine(strategies ...Strategy) *Engine {
	return &Engine{strategies: strategies}
}

func (e *Engine) Run(
	candle models.Candle,
	history []models.Candle,
) []*models.Signal {

	var signals []*models.Signal

	for _, s := range e.strategies {

		sig := s.Evaluate(candle, history)
		if sig != nil {
			signals = append(signals, sig)
		}
	}

	return signals
}
