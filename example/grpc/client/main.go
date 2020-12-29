package main

import (
	"context"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/trace"
	"google.golang.org/grpc"
)

func main() {

	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.Run(testRpc)
}

func testRpc() error {
	// 连接
	tracer, closer, _ := trace.StdConfig().Build()
	conn, _ := grpc.Dial(":8001", grpc.WithInsecure(), trace.ClientDialOption(tracer))

	defer closer.Close()
	defer conn.Close()

	// 初始化客户端
	c := api.NewHelloClient(conn)

	// 调用方法
	req := &api.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		println("err:", err.Error())
	}
	log.Info("success:", res.GetMessage())
	return nil
}
