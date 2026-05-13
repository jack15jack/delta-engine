package portfolio

import (
	"database/sql"
	"time"

	"github.com/jack15jack/delta-engine/internal/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePortfolio(userID string) (*models.Portfolio, error) {

	var id int

	err := s.db.QueryRow(
		`INSERT INTO portfolios (user_id, cash_balance)
		 VALUES ($1, $2)
		 RETURNING id`,
		userID,
		100000,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &models.Portfolio{
		ID:          id,
		UserID:      userID,
		CashBalance: 100000,
	}, nil
}

func (s *Service) GetPortfolio(id int) (*models.Portfolio, error) {

	p := models.Portfolio{}

	err := s.db.QueryRow(
		`SELECT id, user_id, cash_balance
		 FROM portfolios
		 WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.UserID, &p.CashBalance)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Service) GetPositions(portfolioID int) ([]models.Position, error) {

	rows, err := s.db.Query(
		`SELECT id, portfolio_id, ticker, quantity, avg_cost
		 FROM positions
		 WHERE portfolio_id = $1`,
		portfolioID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []models.Position

	for rows.Next() {
		var p models.Position

		err := rows.Scan(
			&p.ID,
			&p.PortfolioID,
			&p.Ticker,
			&p.Quantity,
			&p.AvgCost,
		)

		if err != nil {
			return nil, err
		}

		positions = append(positions, p)
	}

	return positions, nil
}

func (s *Service) ExecuteTrade(portfolioID int, ticker, side string, qty, price float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(
		`INSERT INTO trades (portfolio_id, ticker, side, quantity, price, executed_at)
		 VALUES ($1,$2,$3,$4,$5,$6)`,
		portfolioID,
		ticker,
		side,
		qty,
		price,
		time.Now(),
	)
	if err != nil {
		return err
	}

	// position update
	if side == "BUY" {

		_, err = tx.Exec(`
			INSERT INTO positions (portfolio_id, ticker, quantity, avg_cost)
			VALUES ($1,$2,$3,$4)
			ON CONFLICT (portfolio_id, ticker)
			DO UPDATE SET
				quantity = positions.quantity + EXCLUDED.quantity,
				avg_cost = (
					(positions.avg_cost * positions.quantity + EXCLUDED.avg_cost * EXCLUDED.quantity)
					/ (positions.quantity + EXCLUDED.quantity)
				)
		`, portfolioID, ticker, qty, price)

		if err != nil {
			return err
		}
	}

	if side == "SELL" {

		_, err = tx.Exec(`
			UPDATE positions
			SET quantity = quantity - $3
			WHERE portfolio_id = $1 AND ticker = $2
		`, portfolioID, ticker, qty)

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
