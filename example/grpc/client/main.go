package main

import (
	"context"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// 连接
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := api.NewHelloClient(conn)

	// 调用方法
	req := &api.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		println("err:", err)
	}
	println("success:", res.GetMessage())
}
