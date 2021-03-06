package repositories

import (
	"gorm.io/gorm"
	"task/pkg"
	"task/pkg/models"
	"time"
)

type GiftRepository interface {
	GetGift(code string, phoneNumber string) (*models.Gift, *models.Transaction, error)
	CreateGift(code string, amount int32, batchSize int32) error
}

type GiftRepositoryImpl struct {
	db *pkg.Database
}

func NewGiftRepositoryImpl(db *pkg.Database) *GiftRepositoryImpl {
	return &GiftRepositoryImpl{db: db}
}

func (gr *GiftRepositoryImpl) GetGift(code string, phoneNumber string) (*models.Gift, *models.Transaction, error) {
	var gift models.Gift

	tx := gr.db.DB.Begin()
	result := tx.Table("gift").Where("phone_number is null").Where("code = ?", code).First(&gift).Update("phone_number", phoneNumber)
	if result.Error != nil {
		tx.Rollback()
		return &models.Gift{}, &models.Transaction{}, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return &models.Gift{}, &models.Transaction{}, gorm.ErrRecordNotFound
	}

	var transaction = models.Transaction{
		Amount:      gift.Amount,
		Operation:   int32(models.Increases),
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}

	err := tx.Table("transaction").Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return &models.Gift{}, &models.Transaction{}, err
	}

	tx.Commit()
	return &gift, &transaction, nil
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
