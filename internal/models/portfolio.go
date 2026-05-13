package models

type Portfolio struct {
	ID          int     `json:"id"`
	UserID      string  `json:"user_id"`
	CashBalance float64 `json:"cash_balance"`
}
