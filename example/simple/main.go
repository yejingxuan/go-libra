package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yejingxuan/go-libra/example/simple/api"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/server"
	"github.com/yejingxuan/go-libra/pkg/worker"
	"google.golang.org/grpc"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.AppendServers(httpServer(), grpcServer())
	app.AppendWorkes(weatherWorker())
	app.Run()
}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//V1版本接口定义
	v1 := engine.Group("/service/server/v1/base")
	{
		v1.GET("/healthCheck", func(c *gin.Context) {
			token := c.GetHeader("token")
			cookie, _ := c.Cookie("userType")
			log.Info("token", token)
			/*c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")*/
			log.Info("cookie", cookie)

			rep := gin.H{"message": "ok", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}

//定义grpc-server
func grpcServer() *grpc.Server {
	server := server.GrpcServerStdConfig().Build()
	api.RegisterHelloServer(server, api.HelloService{})
	return server
}

//天气预报任务
func weatherWorker() worker.Worker {
	worker := worker.StdConfig("weather").Build(func() {
		log.Info("任务开始执行,监听天气预报")
	})
	return worker
}
