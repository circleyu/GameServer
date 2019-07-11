package main

import (
	"log"
	"net"

	"./protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 建構體會實作 Calculator 的 gRPC 伺服器。
type server struct{}

func (s *server) SignIn(ctx context.Context, in *protocol.SignInRequest) (*protocol.SignInResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignInResponse{Token: ""}, nil
}

func (s *server) SignOut(ctx context.Context, in *protocol.SignOutRequest) (*protocol.SignOutResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignOutResponse{}, nil
}

func main() {
	// 監聽指定埠口，這樣服務才能在該埠口執行。
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("無法監聽該埠口：%v", err)
	}

	// 建立新 gRPC 伺服器並註冊 Calculator 服務。
	s := grpc.NewServer()
	protocol.RegisterLoginControllerServer(s, &server{})

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(s)

	log.Print("Service Start !!!")

	// 開始在指定埠口中服務。
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法提供服務：%v", err)
	}
}
