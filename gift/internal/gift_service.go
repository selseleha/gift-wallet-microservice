package internal

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"task/gift/api/proto/src"
	"task/pkg"
	"task/pkg/models"
	"task/pkg/repositories"
	walletSrc "task/wallet/api/proto/src"
)

type GiftService struct {
	giftRepo     repositories.GiftRepository
	walletClient walletSrc.WalletServiceClient
	redisService *pkg.RedisService
}

func NewGiftService(giftRepo repositories.GiftRepository, walletClient walletSrc.WalletServiceClient, redisService *pkg.RedisService) *GiftService {
	return &GiftService{
		giftRepo:     giftRepo,
		walletClient: walletClient,
		redisService: redisService,
	}
}

func (gs GiftService) GetGift(ctx context.Context, req *src.GetGiftRequest) (*src.GetGiftResponse, error) {
	gift, transaction, err := gs.giftRepo.GetGift(req.Code, req.PhoneNumber)
	if err != nil {
		return &src.GetGiftResponse{}, err
	}

	walletResponse, err := gs.walletClient.UpdateWallet(ctx, &walletSrc.UpdateWalletRequest{
		PhoneNumber:   req.PhoneNumber,
		Amount:        gift.Amount,
		OperationType: int32(models.Increases),
	})

	if err != nil {
		return &src.GetGiftResponse{}, err
	}

	transactionJson, _ := json.Marshal(transaction)
	gs.redisService.Set(fmt.Sprintf("transaction_%d", transaction.Id), string(transactionJson))

	return &src.GetGiftResponse{
		GiftId:        gift.Id,
		TransactionId: transaction.Id,
		Amount:        walletResponse.LastAmount,
	}, err
}

func (gs GiftService) CreateGift(ctx context.Context, req *src.CreateGiftRequest) (*src.CreateGiftResponse, error) {
	err := gs.giftRepo.CreateGift(req.Code, req.Amount, req.BatchSize)
	return &src.CreateGiftResponse{}, err
}
