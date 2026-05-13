package models

import "time"

type SignalType string

const (
	BUY  SignalType = "BUY"
	SELL SignalType = "SELL"
	HOLD SignalType = "HOLD"
)

type Signal struct {
	Ticker   string     `json:"ticker"`
	Type     SignalType `json:"type"`
	Price    float64    `json:"price"`
	Time     time.Time  `json:"time"`
	Strategy string     `json:"strategy"`
	Strength float64    `json:"strength"` // confidence score (0–1)
}
