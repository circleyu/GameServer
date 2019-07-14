package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	protocol "github.com/CircleYu/GameServer/LoginService/protocol"
	grpclb "github.com/CircleYu/GameServer/etcdv3"
	setting "github.com/CircleYu/GameServer/setting"
	token "github.com/CircleYu/GameServer/token"
	"google.golang.org/grpc"
)

func main() {

	setting.Init()
	token.Init(setting.HmacKeyPath())

	lis, err := net.Listen("tcp", net.JoinHostPort(setting.Host(), setting.Port()))
	if err != nil {
		panic(err)
	}

	err = grpclb.Register(setting.RegisterServer(), setting.ServiceName(), setting.Host(), setting.Port(), time.Second*10, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		grpclb.UnRegister()
		os.Exit(1)
	}()

	log.Printf("starting login service at %s", setting.Port())
	s := grpc.NewServer()
	protocol.RegisterLoginControllerServer(s, &server{})
	s.Serve(lis)
}

type server struct{}

func (s *server) SignIn(ctx context.Context, in *protocol.SignInRequest) (*protocol.SignInResponse, error) {

	// TODO:先驗證帳號密碼是否正確
	// TODO:再去Redis中檢查是否有對應的玩家資料，如果沒有則去DB中撈出來存入Redis中
	// TODO:產生下一次用的token，並更新Redis內容

	log.Printf("Login Account：%v", in.Account)

	tokenString, err := token.CreateToken(in.Account)

	log.Printf("Token：%v", tokenString)

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignInResponse{Token: tokenString}, err
}

func (s *server) SignOut(ctx context.Context, in *protocol.SignOutRequest) (*protocol.SignOutResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignOutResponse{}, nil
}
