# Golang微服务快速构建框架


## quick-start
```go
package main

import (
	"github.com/gin-gonic/gin"
	libra "github.com/yejingxuan/go-libra/pkg"
)

func main() {
    //初始化一个默认的app
	app := libra.DefaultApplication()
    //执行start就绪工作
	app.Start()
	//把自定义server添加到启动server中
	app.AppendServers(httpServer())
    //启动app
	app.Run()

}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
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
```