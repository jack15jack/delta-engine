package models

type Portfolio struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	UserID      string  `json:"user_id"`
	CashBalance float64 `json:"cash_balance"`
}
