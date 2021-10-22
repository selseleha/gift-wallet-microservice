package internal

import (
	"golang.org/x/net/context"
	"task/gift/api/proto/src"
	"task/pkg/repositories"
)

type GiftService struct {
	giftRepo repositories.GiftRepository
}

func NewGiftService(giftRepo repositories.GiftRepository) *GiftService {
	return &GiftService{
		giftRepo: giftRepo,
	}
}

func (gs GiftService) GetGift(ctx context.Context, req *src.GetGiftRequest) (*src.GetGiftResponse, error) {
	gift, err := gs.giftRepo.GetGift(req.Code, req.PhoneNumber)

	return &src.GetGiftResponse{
		GiftId:        gift.Id,
		TransactionId: "",
		Amount:        0,
	}, err
}
