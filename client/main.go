package main

import (
	"flag"
	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
)

func main() {
	//IPアドレス(ここではlocalhost)とポート番号(ここでは5000)を指定して、サーバーと接続する
	conn, err := grpc.Dial(*addrFlag ,grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	//接続は最後に必ず閉じる
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	//サーバーに対してリクエストを送信する
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "wdgk"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}