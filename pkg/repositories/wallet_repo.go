package repositories

import (
	"gorm.io/gorm"
	"task/pkg"
	"task/pkg/models"
)

type WalletRepository interface {
	GetWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error)
	CreateWallet(phoneNumber string, amount int32) error
	UpdateWallet(phoneNumber string, amount int32, operationType int32) error
}

type WalletRepositoryImpl struct {
	db *pkg.Database
}

func NewWalletRepositoryImpl(db *pkg.Database) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{db: db}
}

func (wr *WalletRepositoryImpl) GetWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := wr.db.DB.Table("wallet").Where("phone_number=?", phoneNumber).Find(&wallet).Error; err != nil {
		return &models.Wallet{}, err

	}
	if wallet.Id == 0 {
		return &models.Wallet{}, gorm.ErrRecordNotFound
	}
	return &wallet, nil
}

func (wr *WalletRepositoryImpl) CreateWallet(phoneNumber string, amount int32) error {
	panic("used in fake repo")
}

func (wr *WalletRepositoryImpl) UpdateWallet(phoneNumber string, amount int32, operationType int32) error {

	wallet, err := wr.GetWalletByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}
	if operationType == int32(models.Decreases) {
		err = wr.db.DB.Model(&wallet).Where("id = ?", wallet.Id).Update("amount", wallet.Amount-amount).Error
	}
	if operationType == int32(models.Increases) {
		err = wr.db.DB.Model(&wallet).Where("id = ?", wallet.Id).Update("amount", wallet.Amount+amount).Error
	}
	return err
}
