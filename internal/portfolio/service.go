package portfolio

import (
	"fmt"
	"time"

	"github.com/jack15jack/delta-engine/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreatePortfolio(userID string) (*models.Portfolio, error) {

	p := models.Portfolio{
		UserID:      userID,
		CashBalance: 100000,
	}

	err := s.db.Create(&p).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Service) GetPortfolio(id int) (*models.Portfolio, error) {

	var p models.Portfolio

	err := s.db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Service) GetPositions(portfolioID int) ([]models.Position, error) {

	var positions []models.Position

	err := s.db.Where("portfolio_id = ?", portfolioID).Find(&positions).Error
	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (s *Service) GetTrades(portfolioID int) ([]models.Trade, error) {

	var trades []models.Trade

	err := s.db.Where("portfolio_id = ?", portfolioID).Order("executed_at DESC").Find(&trades).Error

	if err != nil {
		return nil, err
	}

	return trades, nil
}

func (s *Service) ExecuteTrade(portfolioID int, ticker, side string, qty, price float64) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		// 1. Load portfolio (needed for cash updates)
		var portfolio models.Portfolio
		if err := tx.First(&portfolio, portfolioID).Error; err != nil {
			return err
		}

		// 2. Insert trade
		trade := models.Trade{
			PortfolioID: portfolioID,
			Ticker:      ticker,
			Side:        side,
			Quantity:    qty,
			Price:       price,
			ExecutedAt:  time.Now(),
		}

		if err := tx.Create(&trade).Error; err != nil {
			return err
		}

		// 3. Calculate trade value
		tradeValue := qty * price

		// BUY LOGIC
		if side == "BUY" {

			// risk check (basic)
			if portfolio.CashBalance < tradeValue {
				return fmt.Errorf("insufficient cash")
			}

			// update cash
			portfolio.CashBalance -= tradeValue

			// update position
			var pos models.Position
			err := tx.Where("portfolio_id = ? AND ticker = ?", portfolioID, ticker).
				First(&pos).Error

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					pos = models.Position{
						PortfolioID: portfolioID,
						Ticker:      ticker,
						Quantity:    qty,
						AvgCost:     price,
					}

					if err := tx.Create(&pos).Error; err != nil {
						return err
					}
				} else {
					return err
				}
			} else {
				totalQty := pos.Quantity + qty

				pos.AvgCost =
					(pos.AvgCost*pos.Quantity + price*qty) / totalQty

				pos.Quantity = totalQty

				if err := tx.Save(&pos).Error; err != nil {
					return err
				}
			}
		}

		// SELL LOGIC
		if side == "SELL" {

			// update cash
			portfolio.CashBalance += tradeValue

			// update position
			var pos models.Position
			err := tx.Where("portfolio_id = ? AND ticker = ?", portfolioID, ticker).
				First(&pos).Error

			if err != nil {
				return err
			}

			// prevent invalid sells
			if pos.Quantity < qty {
				return fmt.Errorf("insufficient position size")
			}

			pos.Quantity -= qty

			// optional: clear position if zero
			if pos.Quantity == 0 {
				if err := tx.Delete(&pos).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Save(&pos).Error; err != nil {
					return err
				}
			}
		}

		// 4. Persist portfolio update
		if err := tx.Save(&portfolio).Error; err != nil {
			return err
		}

		return nil
	})
}
