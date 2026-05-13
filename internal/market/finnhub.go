package market

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jack15jack/delta-engine/internal/models"

	"github.com/go-resty/resty/v2"
)

type FinnhubProvider struct {
	client *resty.Client
	apiKey string
}

func NewFinnhubProvider() *FinnhubProvider {
	return &FinnhubProvider{
		client: resty.New(),
		apiKey: os.Getenv("FINNHUB_API_KEY"),
	}
}

func (f *FinnhubProvider) GetQuote(symbol string) (*models.Candle, error) {

	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, f.apiKey)

	resp, err := f.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	var data struct {
		C float64 `json:"c"` // close
		H float64 `json:"h"` // high
		L float64 `json:"l"` // low
		O float64 `json:"o"` // open
		T int64   `json:"t"` // time
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}

	return &models.Candle{
		Ticker: symbol,
		Time:   time.Unix(data.T, 0),

		Open:  data.O,
		High:  data.H,
		Low:   data.L,
		Close: data.C,
	}, nil
}

func (f *FinnhubProvider) GetHistory(symbol string, limit int) ([]models.Candle, error) {

	url := fmt.Sprintf(
		"https://finnhub.io/api/v1/stock/candle?symbol=%s&resolution=1&count=%d&token=%s",
		symbol,
		limit,
		f.apiKey,
	)

	resp, err := f.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	var raw struct {
		T []int64   `json:"t"` // time
		O []float64 `json:"o"` // open
		H []float64 `json:"h"` // high
		L []float64 `json:"l"` // low
		C []float64 `json:"c"` // close
		V []float64 `json:"v"` // volume
	}

	err = json.Unmarshal(resp.Body(), &raw)
	if err != nil {
		return nil, err
	}

	var candles []models.Candle

	for i := range raw.T {
		candles = append(candles, models.Candle{
			Ticker: symbol,
			Time:   time.Unix(raw.T[i], 0),

			Open:  raw.O[i],
			High:  raw.H[i],
			Low:   raw.L[i],
			Close: raw.C[i],

			Volume: raw.V[i],
		})
	}

	return candles, nil
}
