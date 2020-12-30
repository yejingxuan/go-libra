package main

import (
	search_api "github.com/yejingxuan/go-libra/example/search/server"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.AppendServers(grpcServer())
	app.Run()
}

//定义grpc-server
func grpcServer() *grpc.Server {
	server := server.GrpcServerStdConfig().Build()
	search_api.RegisterSearchServer(server, search_api.SearchService{})
	return server
}
