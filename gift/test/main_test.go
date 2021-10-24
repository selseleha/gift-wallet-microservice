package test

import (
	"google.golang.org/grpc"
	"os"
	"task/pkg"
	"task/pkg/test"
	walletSrc "task/wallet/api/proto/src"
	test2 "task/wallet/test"
	"testing"
	"time"
)

var fakeRepo = test.NewFakeRepo()
var walletClient walletSrc.WalletServiceClient
var redisConn *pkg.RedisService

func TestMain(m *testing.M) {
	test2.StartWalletMock("0.0.0.0:3003")
	walletConn, _ := grpc.Dial("0.0.0.0:3003", grpc.WithInsecure())
	walletClient = walletSrc.NewWalletServiceClient(walletConn)

	redisOption := pkg.RedisOption{RedisURL: "127.0.0.1:6379"}
	redisConn = pkg.NewRedis(&redisOption)
	time.Sleep(time.Second)
	os.Exit(m.Run())

}
