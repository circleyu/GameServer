package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	protocol "github.com/CircleYu/GameServer/LoginService/protocol"
	grpclb "github.com/CircleYu/GameServer/etcdv3"
	setting "github.com/CircleYu/GameServer/setting"
	token "github.com/CircleYu/GameServer/token"
	"github.com/apsdehal/go-logger"
	"google.golang.org/grpc"
)

var log *logger.Logger
var closeFunc func()

func init() {

	var err error
	if setting.Log2Files() {
		file, err := os.OpenFile(setting.GetLogPath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		closeFunc = func() {
			file.Close()
		}
		log, err = logger.New("LoginService", 0, file)
	} else {
		log, err = logger.New("LoginService", 1)
	}

	if err != nil {
		panic(err) // Check for error
	}

	// Show warning with format message
	log.SetFormat("[%{time}] [%{level}] [%{module}] %{message} (in %{filename}:%{line})")
}

func main() {

	defer closeFunc()

	lis, err := net.Listen("tcp", net.JoinHostPort(setting.Host(), setting.Port()))
	if err != nil {
		log.PanicF("%v", err)
	}

	err = grpclb.Register(setting.RegisterServer(), setting.ServiceName(), setting.Host(), setting.Port(), time.Second*10, 15)
	if err != nil {
		log.FatalF("%v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Infof("receive signal '%v'", s)
		grpclb.UnRegister()
		os.Exit(1)
	}()

	log.Infof("starting login service at %s", setting.Port())

	if setting.Log2Files() {
		fmt.Printf("starting login service at %s", setting.Port())
	}

	s := grpc.NewServer()
	protocol.RegisterLoginControllerServer(s, &server{})
	s.Serve(lis)
}

type server struct{}

type signInData struct {
	Account  string
	Password string
}

type signInReturnData struct {
	Token string
}

func (s *server) SignIn(ctx context.Context, in *protocol.SignInRequest) (*protocol.SignInResponse, error) {

	var data signInData
	err := json.Unmarshal(in.Data, &data)
	if err != nil {
		log.ErrorF("error:%v", err)
		return &protocol.SignInResponse{Data: nil}, err
	}

	log.Infof("Login Account：%v", data.Account)

	// TODO:先驗證帳號密碼是否正確
	// TODO:再去Redis中檢查是否有對應的玩家資料，如果沒有則去DB中撈出來存入Redis中
	// TODO:產生下一次用的token，並更新Redis內容

	tokenString, err := token.CreateToken(data.Account)

	log.Infof("Token：%v", tokenString)

	returnData := signInReturnData{
		Token: tokenString,
	}

	returnBytes, err := json.Marshal(returnData)
	if err != nil {
		log.ErrorF("error:%v", err)
		return &protocol.SignInResponse{}, err
	}
	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignInResponse{Data: returnBytes}, err
}

func (s *server) SignOut(ctx context.Context, in *protocol.SignOutRequest) (*protocol.SignOutResponse, error) {

	// 包裝成 Protobuf 建構體並回傳。
	return &protocol.SignOutResponse{}, nil
}
