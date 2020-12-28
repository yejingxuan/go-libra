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
	return resp, nil
}
