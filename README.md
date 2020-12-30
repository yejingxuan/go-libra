# Golang微服务集成框架

- [Golang微服务集成框架](#golang微服务集成框架)
	- [一、简介](#一简介)
	- [二、quick-start](#二quick-start)
	- [三、功能说明](#三功能说明)
		- [3.1、配置](#31配置)
		- [3.2、日志](#32日志)
		- [3.3、http服务](#33http服务)
		- [3.4、grpc服务](#34grpc服务)
		- [3.5、任务服务](#35任务服务)
		- [3.6、链路追踪](#36链路追踪)

## 一、简介

Golang开发快速集成框架，主要功能
- [x] 配置——viper
- [x] 日志——zap
- [x] http服务——gin
- [x] rpc服务——grpc
- [x] 任务服务——robfig/cron
- [x] 链路追踪——opentracing+jaeger


## 二、quick-start
```go
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
			rep := gin.H{"message": "ok", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}

//定义grpc-server
func grpcServer() *grpc.Server {
	server := server.GrpcStdConfig().Build()
	api.RegisterHelloServer(server, hello.HelloService{})
	return server
}

//天气预报任务
func weatherWorker() worker.Worker {
	worker := worker.StdConfig("weather").Build(func() {
		log.Info("任务开始执行,监听天气预报")
	})
	return worker
}
```

## 三、功能说明

### 3.1、配置

### 3.2、日志

### 3.3、http服务

### 3.4、grpc服务
- grpc服务安装

- grpc的proto文件编写

- grpc的proto生成go编码
	```shell script
	protoc -I . --go_out=plugins=grpc:. ./hello.proto
	```
- 开启grpc服务

### 3.5、任务服务


### 3.6、链路追踪
- 采用jaeger + opentracing 的方式来实现
- jaeger快速搭建
	```shell
	docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778  -p 16686:16686 -p 14268:14268  -p 14269:14269   -p 9411:9411 jaegertracing/all-in-one:latest
	```
- 可视化页面查询链路
  - 访问地址 http://localhost:16686
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20201230151609.png)