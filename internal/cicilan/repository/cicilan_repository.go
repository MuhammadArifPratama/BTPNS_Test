package repository

import (
	"btpntest/domain"

	"gorm.io/gorm"
)

type CicilanRepository struct {
	db *gorm.DB
}

func NewCicilanRepository(db *gorm.DB) *CicilanRepository {
	return &CicilanRepository{db: db}
}

func (r *CicilanRepository) GetAllTenors() ([]domain.Tenor, error) {
	var tenors []domain.Tenor
	if err := r.db.Find(&tenors).Error; err != nil {
		return nil, err
	}
	return tenors, nil
}
