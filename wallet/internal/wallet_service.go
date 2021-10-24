package internal

import (
	"golang.org/x/net/context"
	"task/pkg/repositories"
	"task/wallet/api/proto/src"
)

type WalletService struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (ws WalletService) GetWallet(ctx context.Context, req *src.GetWalletRequest) (*src.GetWalletResponse, error) {
	wallet, err := ws.walletRepo.GetWalletByPhoneNumber(req.PhoneNumber)
	return &src.GetWalletResponse{
		Id:          wallet.Id,
		PhoneNumber: wallet.PhoneNumber,
		Amount:      wallet.Amount,
	}, err
}

func (ws WalletService) UpdateWallet(ctx context.Context, req *src.UpdateWalletRequest) (*src.UpdateWalletResponse, error) {
	err := ws.walletRepo.UpdateWallet(req.PhoneNumber, req.Amount, req.OperationType)
	if err != nil {
		return &src.UpdateWalletResponse{}, err
	}
	wallet, err := ws.walletRepo.GetWalletByPhoneNumber(req.PhoneNumber)
	return &src.UpdateWalletResponse{
		LastAmount: wallet.Amount,
	}, err
}

func (ws WalletService) CreateWallet(ctx context.Context, req *src.CreateWalletRequest) (*src.CreateWalletResponse, error) {
	wallet, err := ws.walletRepo.CreateWallet(req.PhoneNumber, 0)
	return &src.CreateWalletResponse{
		Id:          wallet.Id,
		PhoneNumber: wallet.PhoneNumber,
		Amount:      wallet.Amount,
	}, err
}
