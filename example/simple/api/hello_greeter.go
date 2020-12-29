package hello

import (
	"context"
	"fmt"
	"github.com/yejingxuan/go-libra/example/grpc/client/api"
	"github.com/yejingxuan/go-libra/pkg/log"
)

// 定义helloService并实现约定的接口
type HelloService struct{}

// SayHello 实现Hello服务接口
func (h HelloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Info("hello request: %s", in.Name)
	resp := new(api.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	callSearch(in.Name)
	return resp, nil
}

func callSearch(key string) {
	/*tracer := trace.StdConfig().Build()
	conn, err := grpc.Dial(":8002", grpc.WithInsecure(), trace.ClientDialOption(tracer))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := search_api.NewSearchClient(conn)

	// 调用方法
	req := &search_api.FTRSearchReq{Key: key}
	res, err := c.FTRSearch(context.Background(), req)

	if err != nil {
		println("err:", err.Error())
	}
	println("success:", res.GetResult())*/
}
