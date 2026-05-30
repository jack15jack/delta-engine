package market

import (
	"time"

	"github.com/jack15jack/delta-engine/internal/models"
)

type Provider interface {
	GetQuote(symbol string) (*models.Candle, error)
	GetHistoricalData(symbol string, start time.Time, end time.Time) ([]models.Candle, error)
}
