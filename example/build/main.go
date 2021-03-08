package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yejingxuan/go-libra/example/build/utils"
	libra "github.com/yejingxuan/go-libra/pkg"
)

// 自动发布
func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.AppendServers(httpServer())
	app.Run()
}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//V1版本接口定义
	v1 := engine.Group("/service/v1")
	{
		v1.GET("/build", func(c *gin.Context) {
			utils.DockerBuild("test1", "D:\\dev-source\\go-path\\src\\go-libra\\example\\build\\data\\test-v1.0\\Dockerfile")
			rep := gin.H{"msg": "build-success", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}
