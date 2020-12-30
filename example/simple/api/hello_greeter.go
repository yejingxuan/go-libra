package hello

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	search_api "github.com/yejingxuan/go-libra/example/search/server"
	"github.com/yejingxuan/go-libra/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// 定义helloService并实现约定的接口
type HelloService struct{}

// SayHello 实现Hello服务接口
func (h HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Info("hello request: %s", in.Name)
	resp := new(api.HelloResponse)
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
	tracer := opentracing.GlobalTracer()

	conn, err := grpc.Dial(":8002", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)))
	//conn, err := grpc.DialContext(ctx, ":8002", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		grpclog.Fatalln(err)
	}
	//defer closer.Close()
	defer conn.Close()
	// 初始化客户端
	c := search_api.NewSearchClient(conn)

	// 调用方法
	req := &search_api.FTRSearchReq{Key: key}
	res, err := c.FTRSearch(ctx, req)

	if err != nil {
		println("err:", err.Error())
	}
	println("success:", res.GetResult())

}
