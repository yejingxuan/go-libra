package main

import (
	"context"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/server"
)

func main() {

	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.Run(testRpc)
}

func testRpc() error {
	/*tracer, closer, _ := trace.StdConfig().Build()
	conn, _ := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)))*/
	//初始化客户端
	conn, closer := server.GrpcClientStdConfig("hello").Build()
	c := api.NewHelloClient(conn)
	// 调用方法
	res, err := c.SayHello(context.Background(), &api.HelloRequest{Name: "gRPC"})
	if err != nil {
		log.Error("sayhello调用报错", err)
	}
	log.Info("success:", res.GetMessage())

	if closer != nil {
		defer closer.Close()
	}
	defer conn.Close()
	return nil
}
