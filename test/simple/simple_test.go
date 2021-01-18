package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

//单元测试————simple
func TestGovStatistic(t *testing.T) {
	app := libra.DefaultApplication()
	app.Start()
	server := httpServer()
	//把自定义server添加到启动server中
	//app.AppendServers(server)
	//app.Run()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/service/server/v1/base/healthCheck", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Logger.Info("", zap.Any("w.Body.String()", w.Body.String()))
}

//性能测试————simple
func BenchmarkGovStatistic(b *testing.B) {
	app := libra.DefaultApplication()
	app.Start()
	server := httpServer()

	w := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/service/server/v1/base/healthCheck", nil)
		server.ServeHTTP(w, req)
		//log.Logger.Info("", zap.Any("w.Body.String()", w.Body.String()))
	}
}

//定义http-server
func httpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//V1版本接口定义
	v1 := engine.Group("/service/server/v1/base")
	{
		v1.GET("/healthCheck", func(c *gin.Context) {
			//time.Sleep(2000 * time.Millisecond)
			rep := gin.H{"message": "ok", "code": 200}
			c.JSON(200, rep)
		})
	}
	return engine
}
