package server

import (
	"context"
	"fmt"

	"github.com/yejingxuan/go-libra/pkg/log"
)

// 定义helloService并实现约定的接口
type SearchService struct{}

// SayHello 实现Hello服务接口
func (h SearchService) FTRSearch(ctx context.Context, in *FTRSearchReq) (*FTRSearchRep, error) {
	log.Info("hello request: %s", in.Key)
	resp := new(FTRSearchRep)
	resp.Result = fmt.Sprintf("检索参数： %s.", in.Key)
	return resp, nil
}
