package test

import (
	"google.golang.org/grpc"
	"os"
	"task/pkg/test"
	walletSrc "task/wallet/api/proto/src"
	test2 "task/wallet/test"
	"testing"
)

var fakeRepo = test.NewFakeRepo()
var walletClient walletSrc.WalletServiceClient

func TestMain(m *testing.M) {
	test2.StartWalletMock("0.0.0.0:3003")
	walletConn, _ := grpc.Dial("0.0.0.0:3003", grpc.WithInsecure())
	walletClient = walletSrc.NewWalletServiceClient(walletConn)
	os.Exit(m.Run())

}
