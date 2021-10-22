package repositories

import (
	"task/pkg"
	"task/pkg/models"
)

type GiftRepository interface {
	GetGift(phoneNumber string, code string) (*models.Gift, error)
}

type GiftRepositoryImpl struct {
	db *pkg.Database
}

func NewGiftRepositoryImpl(db *pkg.Database) *GiftRepositoryImpl {
	return &GiftRepositoryImpl{db: db}
}

func (gr *GiftRepositoryImpl) GetGift(phoneNumber string, code string) (*models.Gift, error) {

	return &models.Gift{}, nil
}
