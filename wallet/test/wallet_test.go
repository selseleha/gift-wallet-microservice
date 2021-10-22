package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"task/pkg/models"
	"task/wallet/api/proto/src"
	"task/wallet/internal"
	"testing"
)

const samplePhoneNumber = "09333333333"
const sampleAmount = 1000

func TestWallet(t *testing.T) {

	walletService := internal.NewWalletService(fakeRepo)
	ctx := context.Background()
	fakeRepo.CreateWallet(samplePhoneNumber, sampleAmount)

	t.Run("test get wallet by phone number", func(t *testing.T) {
		wallet, err := walletService.GetWallet(ctx, &src.GetWalletRequest{PhoneNumber: samplePhoneNumber})
		assert.Equal(t, wallet.PhoneNumber, samplePhoneNumber)
		assert.Nil(t, err)
	})

	t.Run("test get wallet by wrong phone number", func(t *testing.T) {
		_, err := walletService.GetWallet(ctx, &src.GetWalletRequest{PhoneNumber: randomPhoneNumber()})
		assert.NotNil(t, err)
	})

	t.Run("test update wallet increases amount", func(t *testing.T) {
		amount := rand.Int31()
		wallet, err := walletService.UpdateWallet(ctx, &src.UpdateWalletRequest{
			PhoneNumber:   samplePhoneNumber,
			Amount:        amount,
			OperationType: int32(models.Increases),
		})
		assert.Equal(t, wallet.LastAmount, amount+sampleAmount)
		assert.Nil(t, err)
	})

	t.Run("test update wrong wallet", func(t *testing.T) {
		amount := rand.Int31()
		_, err := walletService.UpdateWallet(ctx, &src.UpdateWalletRequest{
			PhoneNumber:   randomPhoneNumber(),
			Amount:        amount,
			OperationType: 0,
		})
		assert.NotNil(t, err)
	})

}
