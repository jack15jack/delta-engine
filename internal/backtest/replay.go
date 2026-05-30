package backtest

import "github.com/jack15jack/delta-engine/internal/models"

type Replay struct {
	candles []models.Candle
	index   int
}

func NewReplay(candles []models.Candle) *Replay {
	return &Replay{
		candles: candles,
		index:   0,
	}
}

func (r *Replay) Next() (*models.Candle, bool) {

	if r.index >= len(r.candles) {
		return nil, false
	}

	candle := r.candles[r.index]

	r.index++

	return &candle, true
}
