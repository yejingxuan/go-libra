package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	libra "github.com/yejingxuan/go-libra/pkg"
	"net/http"
	"time"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.AppendServers(httpServer())
	app.Run()
}

func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type",
			"X-Requested-With", "X-Request-ID", "X-HTTP-Method-Override",
			"Content-Type", "Upload-Length", "Upload-Offset", "Tus-Resumable",
			"Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		ExposeHeaders: []string{"Content-Type", "Upload-Offset", "Location",
			"Upload-Length", "Tus-Version", "Tus-Resumable", "Tus-Max-Size",
			"Tus-Extension", "Upload-Metadata", "Upload-Defer-Length", "Upload-Concat"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	}))
	//V1版本接口定义
	v1 := server.Group("/v1")
	{
		v1.GET("/getData", getDataByChunked)
	}
	v2 := server.Group("/v2")
	{
		v2.Use(gzip.Gzip(gzip.DefaultCompression))
		v2.GET("/getData", getData)
	}
	return server
}

func getData(ctx *gin.Context) {
	res := make([]Rep, 0)
	for i := 0; i < 500000; i++ {
		res = append(res, Rep{
			Id:   fmt.Sprintf("158958-zrzhczt_ggfwss_xx_%d", i),
			Lon:  119.34860903,
			Lat:  35.3107758,
			Type: "school",
		})
	}
	ctx.AbortWithStatusJSON(200, res)
}

func getDataByChunked(ctx *gin.Context) {
	w := ctx.Writer
	header := w.Header()
	//在响应头添加分块传输的头字段Transfer-Encoding: chunked
	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := make([]Rep, 0)
	res = append(res, Rep{
		Id:   fmt.Sprintf("158958-zrzhczt_ggfwss_xx_"),
		Lon:  119.34860903,
		Lat:  35.3107758,
		Type: "school",
	})

	//Flush()方法，好比服务端在往一个文件中写了数据，浏览器会看见此文件的内容在不断地增加。
	w.Write([]byte(fmt.Sprintf("%s", res)))
	w.(http.Flusher).Flush()

	for i := 0; i < 10; i++ {
		w.Write([]byte(fmt.Sprintf(`hello %d`, i)))
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(100) * time.Millisecond)
	}

	/*w.Write([]byte(`hello-end`))
	w.(http.Flusher).Flush()*/
}

type Rep struct {
	Id   string
	Lon  float32
	Lat  float32
	Type string
}
