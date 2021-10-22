package api

import (
	"golang.org/x/net/context"
	"task/wallet/api/proto/src"
	"task/wallet/internal"
)

type WalletHandlerImpl struct {
	WalletService *internal.WalletService
}

func NewWalletHandlerImpl(walletService *internal.WalletService) *WalletHandlerImpl {
	return &WalletHandlerImpl{
		WalletService: walletService,
	}
}

func (w WalletHandlerImpl) GetWallet(ctx context.Context, req *src.GetWalletRequest) (*src.GetWalletResponse, error) {
	return w.WalletService.GetWallet(ctx, req)
}

func (w WalletHandlerImpl) UpdateWallet(ctx context.Context, req *src.UpdateWalletRequest) (*src.UpdateWalletResponse, error) {

	return w.WalletService.UpdateWallet(ctx, req)

}
