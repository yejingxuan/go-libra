package main

import (
	"github.com/gin-gonic/gin"
	libra "github.com/yejingxuan/go-libra/pkg"
)

func main() {
	app := libra.DefaultApplication()
	//把自定义server添加到启动server中
	app.AppendServers(httpServer())
	app.Startup()
}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	//V1版本接口定义
	v1 := engine.Group("/service/api/v1/base")
	{
		v1.GET("/healthCheck", func(c *gin.Context) {
			rep := gin.H{
				"message": "ok",
				"code":    200,
			}
			c.JSON(200, rep)
		})
	}
	return engine
}
