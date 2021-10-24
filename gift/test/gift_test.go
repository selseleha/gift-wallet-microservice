package test

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"task/gift/api/proto/src"
	"task/gift/internal"
	"task/pkg/utils"
	"testing"
)

const samplePhoneNumber = "09333333333"
const sampleAmount = 1000
const sampleCode = "IR_BR"

func TestGift(t *testing.T) {

	giftService := internal.NewGiftService(fakeRepo, walletClient, redisConn)
	ctx := context.Background()
	fakeRepo.CreateGift(sampleCode, sampleAmount, 100)
	fakeRepo.CreateWallet(samplePhoneNumber, 0)

	t.Run("test get gift by phone number and code", func(t *testing.T) {
		gift, err := giftService.GetGift(ctx, &src.GetGiftRequest{
			PhoneNumber: samplePhoneNumber,
			Code:        sampleCode,
		})
		assert.NotEqual(t, gift.GiftId, 0)
		assert.Nil(t, err)
	})

	t.Run("test get gift by phone number and wrong code", func(t *testing.T) {
		gift, err := giftService.GetGift(ctx, &src.GetGiftRequest{
			PhoneNumber: samplePhoneNumber,
			Code:        utils.GenerateRandomString(5),
		})
		assert.Equal(t, gift.GiftId, int32(0))
		assert.NotNil(t, err)
	})
}
