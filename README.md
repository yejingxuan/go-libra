# Golang微服务集成框架

- [Golang微服务集成框架](#golang微服务集成框架)
  - [一、简介](#一简介)
  - [二、quick-start](#二quick-start)
  - [三、功能说明](#三功能说明)
    - [3.1、配置](#31配置)
    - [3.2、日志](#32日志)
    - [3.3、http服务](#33http服务)
    - [3.4、grpc服务](#34grpc服务)
    - [3.5、etcd](#35etcd)
    - [3.5、任务服务](#35任务服务)
    - [3.6、链路追踪](#36链路追踪)
    - [3.7、工具包](#37工具包)
  - [四、规划](#四规划)
    - [4.1、TODO](#41todo)

## 一、简介

Golang开发快速集成框架，主要功能
- [x] 配置——viper
- [x] 日志——zap
- [x] 协程池——ants
- [x] http服务——gin
- [x] rpc服务——grpc
- [x] 任务服务——robfig/cron
- [x] 链路追踪——opentracing+jaeger
- [x] 加解密工具包 


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
- 配置采用viper配置框架来进行集成，并结合项目跟目录下的config.toml来进行辅助配置

### 3.2、日志
- 日志采用zap日志框架来进行日志采集，默认日志输出在项目根目录下的 log 文件夹
- 也可通过配置文件来指定日志的存储路径
  ```toml
  [general]
  log_path = "./log/"
  ```

### 3.3、http服务
- http服务目前提供了开源框架gin来提供

### 3.4、grpc服务
- grpc服务安装
    ```go
    go get github.com/golang/protobuf/proto
    go get google.golang.org/grpc
    go get github.com/golang/protobuf/protoc-gen-go
    //安装好后，会在GOPATH/bin下生成protoc-gen-go.exe
    //手动下载protoc.exe，放在GOPATH/bin下
    //注意添加 GOPATH/bin 到系统环境变量中，能够直接使用bin下的命令行工具
    ```
- grpc的proto文件编写

- grpc的proto生成go编码
    ```go
    protoc -I . --go_out=plugins=grpc:. ./hello.proto
    ```
- 开启grpc服务

- 使用etcd来做grpc的服务注册发现和负载均衡

### 3.5、etcd
- etcd搭建
    ```docker
    # 快速搭建etcd
    docker run  -d -p 2379:2379 -p 2380:2380 -p 4001:4001  -p 7001:7001 -e "ETCD_ADVERTISE_CLIENT_URLS=http://localhost:2379" -e "ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379" -e "ETCD_INITIAL_ADVERTISE_PEER_URLS=http://0.0.0.0:2380" -e "ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380"  -e "ALLOW_NONE_AUTHENTICATION=yes" -e "ETCD_INITIAL_CLUSTER=node1=http://0.0.0.0:2380" -e "ETCD_NAME=node1" --name server-etcd3  bitnami/etcd:3
    # 快速搭建etcd可视化管理工具
    docker run -d -p 9222:8080 evildecay/etcdkeeper
    ```
- golang集成etcd
    - 由于etcd-v3的api和grpc版本有冲突，所有使用etcd-v3版本最高智能使用1.26.0的grpc。在go.mod中进行grpc版本替换
        ```go
        replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
        ```
    - 同时protoc-gen-go的版本也要更换到相对应的grpc版本，并重新生成pb.go文件（这个地方还是贼蛋疼的）
        ```go
        go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
        ```

### 3.5、任务服务
- libra 提供了 corn 的定时任务组件，可以通过简单配置快速启动任务
- 配置
  - 任务默认执行频率为 0/2 * * * * ? 即2秒执行一次
  - 可通过配置文件 config.toml 来快速配置任务相关信息
    ```toml
    [worker.weather]
    name = "weather"
    corn = "0/20 * * * * ?"
    ```
- 快速开始
    ```go
    package main

    import (
        libra "github.com/yejingxuan/go-libra/pkg"
        "github.com/yejingxuan/go-libra/pkg/log"
        "github.com/yejingxuan/go-libra/pkg/worker"
    )

    func main() {
        app := libra.DefaultApplication()
        app.Start()
        app.AppendWorkes(weatherWorker()) //支持传入多个任务并行执行
        app.Run()
    }

    //天气预报任务
    func weatherWorker() worker.Worker {
        worker := worker.StdConfig("weather").Build(func() {
            log.Info("任务开始执行111")
        })
        return worker
    }
    ```

### 3.6、链路追踪
- 采用 jaeger + opentracing 的方式来实现
- docker快速搭建 jaeger
    ```docker
    docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778  -p 16686:16686 -p 14268:14268  -p 14269:14269   -p 9411:9411 jaegertracing/all-in-one:latest
    ```
- 可视化页面查询链路
  - 访问地址 http://localhost:16686
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20201230151609.png)

### 3.7、工具包
- 加解密
  > utils工具包提供了一些主流的加解密工具
  - Bcrypt 加密算法（hash算法加密）
  - AES-CBC 加密算法（对称加密）
  - MD5 加密算法（不可逆加密）

- 时间转换

## 四、规划

### 4.1、TODO
- [ ] 工作流-workflow
- [ ] 负载均衡
- [ ] 路由网关限流
- [ ] 分布式缓存
- [ ] 搜索引擎