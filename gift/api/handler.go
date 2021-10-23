package api

import (
	"golang.org/x/net/context"
	"task/gift/api/proto/src"
	"task/gift/internal"
)

type GiftHandlerImpl struct {
	GiftService *internal.GiftService
}

func NewGiftHandlerImpl(giftService *internal.GiftService) *GiftHandlerImpl {
	return &GiftHandlerImpl{
		GiftService: giftService,
	}
}

func (g GiftHandlerImpl) GetGift(ctx context.Context, req *src.GetGiftRequest) (*src.GetGiftResponse, error) {
	return g.GiftService.GetGift(ctx, req)
}

func (g GiftHandlerImpl) CreateGift(ctx context.Context, req *src.CreateGiftRequest) (*src.CreateGiftResponse, error) {
	return g.GiftService.CreateGift(ctx, req)
}
