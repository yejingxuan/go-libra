package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	hello "github.com/yejingxuan/go-libra/example/simple/api"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/worker"
	"google.golang.org/grpc"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.AppendServers(httpServer(), grpcServer())
	//app.AppendWorkes(weatherWorker(), eatWorker())
	app.Run()
}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//V1版本接口定义
	v1 := engine.Group("/service/api/v1/base")
	{
		v1.GET("/healthCheck", func(c *gin.Context) {
			rep := gin.H{"message": "ok", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}

//定义grpc-server
func grpcServer() *grpc.Server {
	server := grpc.NewServer()
	api.RegisterHelloServer(server, hello.HelloService{})
	return server
}

//天气预报任务
func weatherWorker() worker.Worker {
	worker := worker.StdConfig("weather").Build(func() {
		log.Info("任务开始执行111")
	})
	return worker
}

//团建任务
func eatWorker() worker.Worker {
	worker := worker.StdConfig("eat").Build(func() {
		log.Info("任务开始执行222")
	})
	return worker
}
