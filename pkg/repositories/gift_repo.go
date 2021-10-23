package repositories

import (
	"gorm.io/gorm"
	"task/pkg"
	"task/pkg/models"
)

type GiftRepository interface {
	GetGift(code string, phoneNumber string) (*models.Gift, error)
	CreateGift(code string, amount int32, batchSize int32) error
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
	result := tx.Table("gift").Where("phone_number is null").Where("code = ?", code).First(&gift).Update("phone_number", phoneNumber)
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

func (gr *GiftRepositoryImpl) CreateGift(code string, amount int32, batchSize int32) error {
	var gifts []models.Gift
	for i := 0; i < int(batchSize); i++ {
		gifts = append(gifts, models.Gift{
			Code:   code,
			Amount: amount,
		})
	}
	err := gr.db.DB.Create(&gifts).Error
	return err
}
