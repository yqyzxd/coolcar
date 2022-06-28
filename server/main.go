package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	go startGRPCGateway()

	///////启动grpc服务/////////////////
	//设置grpc地址端口
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	//生成grpc server
	s := grpc.NewServer()
	//注册服务
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	//开启服务
	log.Fatal(s.Serve(lis))
}

//开启grpcgateway
func startGRPCGateway() {
	//ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	//mux multiplexer 一对多      参数ServeMuxOption 会在grpc转 restful 或restful grpc时执行   JSONP是默认的Marshaler  这里的处理是将enum转成int
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true, //将enum转成相应的整型
			UseProtoNames:  true, //json 的key 使用_分割的风格
		},
	}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		"localhost:8081", //grpc 地址
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)

	if err != nil {
		log.Fatalf("cannot start grpc gateway:%v", err)
	}

	//gateway 端口 给http请求用的 监听8080端口
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listenr and serve:%v", err)
	}
}
