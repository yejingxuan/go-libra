package main

import (
	"github.com/gin-gonic/gin"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/worker"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	//把自定义server添加到启动server中
	app.AppendServers(httpServer())
	app.AppendWorkes(weatherWorker(), eatWorker())
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

func weatherWorker() worker.Worker {
	worker := worker.StdConfig().Build(func() {
		log.Info("任务开始执行111")
	})
	return worker
}

func eatWorker() worker.Worker {
	config := worker.WorkerConfig{
		WorkerName: "eatWorker",
		WorkerCron: "0/5 * * * * ",
	}
	worker := config.Build(func() {
		log.Info(config.WorkerName + "任务开始执行222")
	})
	return worker
}
