package repositories

import (
	"task/pkg"
	"task/pkg/models"
	"task/pkg/utils"
)

type WalletRepository interface {
	GetWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error)
	CreateWallet(phoneNumber string, amount int32) (*models.Wallet, error)
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
		wallet = models.Wallet{
			PhoneNumber: phoneNumber,
			Amount:      0,
		}
		createWallet, err := wr.CreateWallet(phoneNumber, 0)
		if err != nil {
			return &models.Wallet{}, err
		}
		return createWallet, nil
	}
	return &wallet, nil
}

func (wr *WalletRepositoryImpl) CreateWallet(phoneNumber string, amount int32) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := wr.db.DB.Table("wallet").Where("phone_number=?", phoneNumber).Find(&wallet).Error; err != nil {
		return &models.Wallet{}, err
	}
	if wallet.Id != 0 {
		return &wallet, utils.WalletExistError
	}
	createWallet := models.Wallet{
		PhoneNumber: phoneNumber,
		Amount:      amount,
	}
	err := wr.db.DB.Create(&createWallet).Error
	return &createWallet, err
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
