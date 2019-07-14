package main

import (
	"context"
	"flag"
	"log"
	"time"

	protocol "github.com/CircleYu/GameServer/LoginService/protocol"
	grpclb "github.com/CircleYu/GameServer/etcdv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var (
	svc = flag.String("service", "login_service", "service name")
	reg = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {

	flag.Parse()
	r := grpclb.NewResolver(*reg, *svc)
	resolver.Register(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	// The gRPC client library will use the specified scheme to pick the right resolver plugin and pass it the fully qualified name string.
	conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+*svc, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	cancel()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// 建立新的 Calculator 客戶端，所以等一下就能夠使用 Calculator 的所有方法。
	c := protocol.NewLoginControllerClient(conn)

	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
	log.Printf("發送請求....")
	tokenResp, err := c.SignIn(context.Background(), &protocol.SignInRequest{Account: "circle", Password: "123456"})
	if err != nil {
		log.Fatalf("無法執行 SignIn 函式：%v", err)
	}
	log.Printf("回傳結果：%v", tokenResp.Token)

	/*
		time.Sleep(16 * 1000 * time.Millisecond)

			// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
			CheckResp, err := c.Check(context.Background(), &protocol.CheckRequest{Token: tokenResp.Token})
			if err != nil {
				log.Fatalf("無法執行 Check 函式：%v", err)
			}
			log.Printf("回傳結果：%v", CheckResp.Result)
	*/
}
