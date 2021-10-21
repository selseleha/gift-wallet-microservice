package internal

import (
	"golang.org/x/net/context"
	"task/wallet/api/proto/src"
)

type WalletService struct {
}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (ws WalletService) GetWallet(ctx context.Context, req *src.GetWalletRequest) (*src.GetWalletResponse, error) {
	return nil, nil
}

func (ws WalletService) UpdateWallet(ctx context.Context, req *src.UpdateWalletRequest) (*src.UpdateWalletResponse, error) {
	return nil, nil
}
