package repos

import (
	"github.com/jack15jack/delta-engine/internal/models"
	"gorm.io/gorm"
)

type SignalRepo struct {
	db *gorm.DB
}

func NewSignalRepo(db *gorm.DB) *SignalRepo {
	return &SignalRepo{db: db}
}

func (r *SignalRepo) Insert(signal models.Signal) error {
	return r.db.Create(&signal).Error
}
