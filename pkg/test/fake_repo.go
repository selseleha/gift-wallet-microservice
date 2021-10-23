package test

import (
	"errors"
	"math/rand"
	"task/pkg/models"
	"task/pkg/utils"
	"time"
)

type FakeRepo struct {
	wallets          []models.Wallet
	gifts            []models.Gift
	transactions     []models.Transaction
	walletIndex      int32
	giftIndex        int32
	transactionIndex int32
}

func (r *FakeRepo) GetWalletByPhoneNumber(phoneNumber string) (*models.Wallet, error) {

	for _, wallet := range r.wallets {
		if wallet.PhoneNumber == phoneNumber {
			return &wallet, nil
		}
	}
	return &models.Wallet{}, errors.New("wallet not found")
}

func (r *FakeRepo) CreateWallet(phoneNumber string, amount int32) error {
	wallet := models.Wallet{
		Id:          rand.Int31(),
		PhoneNumber: phoneNumber,
		Amount:      amount,
	}

	r.wallets = append(r.wallets, wallet)
	r.walletIndex++
	return nil
}

func (r *FakeRepo) UpdateWallet(phoneNumber string, amount int32, operationType int32) error {

	if operationType == int32(models.Decreases) {
		for index, wallet := range r.wallets {
			if wallet.PhoneNumber == phoneNumber {
				updateRequest := &r.wallets[index]
				updateRequest.Amount = wallet.Amount - amount
				return nil
			}
		}
	}
	if operationType == int32(models.Increases) {
		for index, wallet := range r.wallets {
			if wallet.PhoneNumber == phoneNumber {
				updateRequest := &r.wallets[index]
				updateRequest.Amount = wallet.Amount + amount
				return nil
			}
		}
	}
	return errors.New("wallet not found")
}

func (r *FakeRepo) GetGift(code string, phoneNumber string) (*models.Gift, *models.Transaction, error) {
	for index, gift := range r.gifts {
		if gift.Code == code && gift.PhoneNumber == nil {
			updateGift := &r.gifts[index]
			updateGift.PhoneNumber = &phoneNumber
			transaction := models.Transaction{
				Id:          rand.Int31(),
				Amount:      gift.Amount,
				Operation:   int32(models.Increases),
				PhoneNumber: phoneNumber,
				CreatedAt:   time.Now(),
			}
			r.transactions = append(r.transactions, transaction)
			return &gift, &transaction, nil
		}
	}
	return &models.Gift{}, &models.Transaction{}, utils.WalletNouFoundError
}

func (r *FakeRepo) CreateGift(code string, amount int32, batchSize int32) error {
	var n int32 = 1
	for n <= batchSize {
		r.gifts = append(r.gifts, models.Gift{
			Id:     int32(n),
			Code:   code,
			Amount: amount,
		})
		r.giftIndex++
		n++
	}
	return nil
}

func (r *FakeRepo) GetTransactions() ([]models.Transaction, error) {
	return r.transactions, nil
}

func NewFakeRepo() *FakeRepo {
	return &FakeRepo{
		wallets:     nil,
		gifts:       nil,
		walletIndex: 0,
		giftIndex:   0,
	}
}
