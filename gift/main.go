package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task/gift/api"
	"task/gift/api/proto/src"
	"task/gift/internal"
	"task/pkg"
	"task/pkg/models"
	"task/pkg/repositories"
	walletSrc "task/wallet/api/proto/src"
)

func main() {
	dbOption := pkg.Option{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "123456",
		Db:   "discount",
	}
	mysqlConnection := pkg.NewMysql(dbOption)

	giftRepo := repositories.NewGiftRepositoryImpl(mysqlConnection)
	mysqlConnection.DB.AutoMigrate(&models.Gift{})

	path := "0.0.0.0:3001"
	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", path)
	grpcServer := grpc.NewServer()
	walletConn, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	walletClient := walletSrc.NewWalletServiceClient(walletConn)

	giftService := internal.NewGiftService(giftRepo, walletClient)
	handler := api.NewGiftHandlerImpl(giftService)
	src.RegisterGiftServiceServer(grpcServer, handler)

	grpcServer.Serve(lis)

}
