package market

import (
	"github.com/jack15jack/delta-engine/internal/models"
)

type Provider interface {
	GetQuote(symbol string) (*models.Candle, error)
	GetHistory(symbol string, limit int) ([]models.Candle, error)
}
