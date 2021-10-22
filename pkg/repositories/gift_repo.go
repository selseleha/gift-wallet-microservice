package repositories

import (
	"gorm.io/gorm"
	"task/pkg"
	"task/pkg/models"
)

type GiftRepository interface {
	GetGift(code string, phoneNumber string) (*models.Gift, error)
}

type GiftRepositoryImpl struct {
	db *pkg.Database
}

func NewGiftRepositoryImpl(db *pkg.Database) *GiftRepositoryImpl {
	return &GiftRepositoryImpl{db: db}
}

func (gr *GiftRepositoryImpl) GetGift(code string, phoneNumber string) (*models.Gift, error) {
	var gift models.Gift
	tx := gr.db.DB.Begin()
	result := tx.Table("gift").Where("phone_number is null").Where("code = ?", code).Update("phone_number", phoneNumber)
	if result.Error != nil {
		tx.Rollback()
		return &models.Gift{}, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return &models.Gift{}, gorm.ErrRecordNotFound
	}

	tx.Commit()
	return &gift, nil
}
