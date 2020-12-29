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
	v1 := engine.Group("/service/server/v1/base")
	{
		v1.GET("/healthCheck", func(c *gin.Context) {
			rep := gin.H{"message": "ok", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}
```

## grpc
- grpc的pb生成
```shell script
protoc -I . --go_out=plugins=grpc:. ./hello.proto
```

## 链路追踪
采用jaeger + opentracing 的方式来实现
- jaeger快速搭建
```shell
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778  -p 16686:16686 -p 14268:14268  -p 14269:14269   -p 9411:9411 jaegertracing/all-in-one:latest
```
- 页面链路查询
  - 访问地址 http://localhost:16686