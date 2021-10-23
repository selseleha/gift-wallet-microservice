package test

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	walletSrc "task/wallet/api/proto/src"
)

type WalletMockHandlerImpl struct {
}

func (w WalletMockHandlerImpl) GetWallet(ctx context.Context, request *walletSrc.GetWalletRequest) (*walletSrc.GetWalletResponse, error) {
	panic("implement me")
}

func (w WalletMockHandlerImpl) UpdateWallet(ctx context.Context, request *walletSrc.UpdateWalletRequest) (*walletSrc.UpdateWalletResponse, error) {
	return &walletSrc.UpdateWalletResponse{
		TransactionId: "",
		LastAmount:    0,
	}, nil
}

func NewWalletMockHandlerImpl() *WalletMockHandlerImpl {
	return &WalletMockHandlerImpl{}
}

func StartWalletMock(address string) {
	grpcServer := grpc.NewServer()
	//var fakeRepo = test.NewFakeRepo()

	//walletService := internal.NewWalletService(fakeRepo)
	//handler := api.NewWalletHandlerImpl(walletService)
	handler := NewWalletMockHandlerImpl()

	walletSrc.RegisterWalletServiceServer(grpcServer, handler)

	go start(grpcServer, address)
}

func start(server *grpc.Server, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		println(err.Error())
	}
	println("Start listening on address: ", address)
	if err := server.Serve(lis); err != nil {
		println("error in starting grpc server: failed to serve")
	}
}
