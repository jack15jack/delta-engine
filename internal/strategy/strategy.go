package strategy

import "github.com/jack15jack/delta-engine/internal/models"

type Strategy interface {
	Name() string
	Evaluate(candle models.Candle, history []models.Candle) *models.Signal
}
