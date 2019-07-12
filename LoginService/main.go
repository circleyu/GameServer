package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	protocol "github.com/CircleYu/GameServer/LoginService/protocol"
	grpclb "github.com/CircleYu/GameServer/etcdv3"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SignIn(ctx context.Context, in *protocol.SignInRequest) (*protocol.SignInResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignInResponse{Token: ""}, nil
}

func (s *server) SignOut(ctx context.Context, in *protocol.SignOutRequest) (*protocol.SignOutResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignOutResponse{}, nil
}

var (
	serv = flag.String("service", "login_service", "service name")
	host = flag.String("host", "localhost", "listening host")
	port = flag.String("port", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {

	flag.Parse()

	lis, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		panic(err)
	}

	err = grpclb.Register(*reg, *serv, *host, *port, time.Second*10, 15)
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

	log.Printf("starting login service at %s", *port)
	s := grpc.NewServer()
	protocol.RegisterLoginControllerServer(s, &server{})
	s.Serve(lis)
}
