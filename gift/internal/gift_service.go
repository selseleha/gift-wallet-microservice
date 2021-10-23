package internal

import (
	"golang.org/x/net/context"
	"task/gift/api/proto/src"
	"task/pkg/models"
	"task/pkg/repositories"
	walletSrc "task/wallet/api/proto/src"
)

type GiftService struct {
	giftRepo     repositories.GiftRepository
	walletClient walletSrc.WalletServiceClient
}

func NewGiftService(giftRepo repositories.GiftRepository, walletClient walletSrc.WalletServiceClient) *GiftService {
	return &GiftService{
		giftRepo:     giftRepo,
		walletClient: walletClient,
	}
}

func (gs GiftService) GetGift(ctx context.Context, req *src.GetGiftRequest) (*src.GetGiftResponse, error) {
	gift, err := gs.giftRepo.GetGift(req.Code, req.PhoneNumber)
	if err != nil {
		return &src.GetGiftResponse{}, err
	}
	walletResponse, err := gs.walletClient.UpdateWallet(ctx, &walletSrc.UpdateWalletRequest{
		PhoneNumber:   req.PhoneNumber,
		Amount:        gift.Amount,
		OperationType: int32(models.Increases),
	})

	return &src.GetGiftResponse{
		GiftId:        gift.Id,
		TransactionId: "",
		Amount:        walletResponse.LastAmount,
	}, err
}

func (gs GiftService) CreateGift(ctx context.Context, req *src.CreateGiftRequest) (*src.CreateGiftResponse, error) {
	err := gs.giftRepo.CreateGift(req.Code, req.Amount, req.BatchSize)
	return &src.CreateGiftResponse{}, err
}
