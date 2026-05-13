package queries

import (
	"database/sql"

	"github.com/jack15jack/delta-engine/internal/models"
)

type SignalRepo struct {
	db *sql.DB
}

func NewSignalRepo(db *sql.DB) *SignalRepo {
	return &SignalRepo{db: db}
}

func (r *SignalRepo) Insert(signal models.Signal) error {

	query := `
	INSERT INTO signals
	(ticker, strategy, signal_type, price, strength, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		query,
		signal.Ticker,
		signal.Strategy,
		signal.Type,
		signal.Price,
		signal.Strength,
		signal.Time,
	)

	return err
}
