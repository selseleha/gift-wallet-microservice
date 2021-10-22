package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task/pkg"
	"task/pkg/models"
	"task/pkg/repositories"
	"task/wallet/api/proto"
	"task/wallet/api/proto/src"
	"task/wallet/internal"
)

func main() {

	dbOption := pkg.Option{
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "123456",
		Db:   "wallet",
	}
	mysqlConnection := pkg.NewMysql(dbOption)

	walletRepo := repositories.NewWalletRepositoryImpl(mysqlConnection)
	mysqlConnection.DB.AutoMigrate(&models.Wallet{})

	path := "0.0.0.0:3000"
	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", path)
	grpcServer := grpc.NewServer()
	walletService := internal.NewWalletService(walletRepo)
	handler := proto.NewWalletHandlerImpl(walletService)
	src.RegisterWalletServiceServer(grpcServer, handler)
	grpcServer.Serve(lis)
}
