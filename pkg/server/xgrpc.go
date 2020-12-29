package server

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yejingxuan/go-libra/pkg/trace"
	"google.golang.org/grpc"
)

type GrpcConfig struct {
	TraceEnable bool
	Port        int
	Protocol    string
}

func GrpcStdConfig() GrpcConfig {
	return GrpcRawConfig(fmt.Sprintf("server"))
}

func GrpcRawConfig(name string) GrpcConfig {
	config := GrpcConfig{
		TraceEnable: viper.GetBool(fmt.Sprintf("%s.grpc_trace_enable", name)),
		Port:        viper.GetInt(fmt.Sprintf("%s.grpc_port", name)),
	}
	return config
}

func (config GrpcConfig) Build() *grpc.Server {
	if config.TraceEnable {
		//开启链路追踪
		tracer, _, _ := trace.StdConfig().Build()
		server := grpc.NewServer(trace.ServerOption(tracer))
		return server
	}
	return grpc.NewServer()
}
