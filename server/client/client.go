package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//连接服务器
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}
	//创建客户端服务代理
	serviceClient := trippb.NewTripServiceClient(conn)
	//发起请求
	resp, err := serviceClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "456",
	})
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}
	fmt.Println(resp)
}
