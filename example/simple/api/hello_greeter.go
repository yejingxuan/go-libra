package api

import (
	"context"
	"fmt"
	search_api "github.com/yejingxuan/go-libra/example/search/server"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/server"
)

// 定义helloService并实现约定的接口
type HelloService struct{}

// SayHello 实现Hello服务接口
func (h HelloService) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Info("hello request: %s", in.Name)
	resp := new(HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	callSearch(ctx, in.Name)
	return resp, nil
}

func callSearch(ctx context.Context, key string) {
	/*if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		log.Info("parentspan:%s, ctx:%s", parent, ctx)
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			mysqlSpan := tracer.StartSpan("FindUserTable", opentracing.ChildOf(pctx))

			//do mysql operations
			time.Sleep(time.Millisecond * 100)

			defer mysqlSpan.Finish()
		}
	}*/
	//tracer, closer, _ := trace.StdConfig().Build()
	// 初始化客户端
	conn, closer := server.GrpcClientStdConfig("search").Build()
	c := search_api.NewSearchClient(conn)
	// 调用方法
	res, err := c.FTRSearch(ctx, &search_api.FTRSearchReq{Key: key})

	if err != nil {
		log.Error("callSearch调用报错", err)
	}
	log.Info("success:", res.GetResult())

	if closer != nil {
		defer closer.Close()
	}
	defer conn.Close()
}
