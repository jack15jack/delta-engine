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

func (f *FinnhubProvider) GetHistoricalData(symbol string, start time.Time, end time.Time) ([]models.Candle, error) {

	from := start.Unix()
	to := end.Unix()

	url := fmt.Sprintf(
		"https://finnhub.io/api/v1/stock/candle?symbol=%s&resolution=D&from=%d&to=%d&token=%s",
		symbol,
		from,
		to,
		f.apiKey,
	)

	resp, err := f.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	fmt.Println("FINNHUB RAW RESPONSE:", string(resp.Body()))

	if len(resp.Body()) == 0 {
		return nil, fmt.Errorf("empty response from finnhub")
	}

	var raw struct {
		S string    `json:"s"`
		T []int64   `json:"t"`
		O []float64 `json:"o"`
		H []float64 `json:"h"`
		L []float64 `json:"l"`
		C []float64 `json:"c"`
		V []float64 `json:"v"`
	}

	if err := json.Unmarshal(resp.Body(), &raw); err != nil {
		return nil, err
	}

	if raw.S != "ok" {
		return nil, fmt.Errorf("finnhub error status: %s", raw.S)
	}

	candles := make([]models.Candle, 0, len(raw.T))

	for i := range raw.T {
		candles = append(candles, models.Candle{
			Ticker: symbol,
			Time:   time.Unix(raw.T[i], 0),
			Open:   raw.O[i],
			High:   raw.H[i],
			Low:    raw.L[i],
			Close:  raw.C[i],
			Volume: raw.V[i],
		})
	}

	return candles, nil
}
