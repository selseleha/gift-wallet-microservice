package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task/wallet/api/proto"
	"task/wallet/api/proto/src"
	"task/wallet/internal"
)

func main() {

	path := "0.0.0.0:3000"
	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", path)
	grpcServer := grpc.NewServer()
	walletService := internal.NewWalletService()
	handler := proto.NewWalletHandlerImpl(walletService)
	src.RegisterWalletServiceServer(grpcServer, handler)
	grpcServer.Serve(lis)
}
