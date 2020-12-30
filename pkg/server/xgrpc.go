package server

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/trace"
	"google.golang.org/grpc"
	"io"
)

//server配置信息
type GrpcServerConfig struct {
	TraceEnable bool
	Port        int
	Protocol    string
}

func GrpcServerStdConfig() GrpcServerConfig {
	return grpcServerRawConfig(fmt.Sprintf("server"))
}

func grpcServerRawConfig(name string) GrpcServerConfig {
	config := GrpcServerConfig{
		TraceEnable: viper.GetBool("server.grpc_trace_enable"),
		Port:        viper.GetInt(fmt.Sprintf("%s.grpc_port", name)),
	}
	return config
}

func (config GrpcServerConfig) Build() *grpc.Server {
	if config.TraceEnable {
		//开启链路追踪
		tracer, _, _ := trace.StdConfig().Build()
		server := grpc.NewServer(
			grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
			grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)))
		return server
	}
	return grpc.NewServer()
}

//client配置信息
type GrpcClientConfig struct {
	TraceEnable bool
	HostPort    string
}

func GrpcClientStdConfig(name string) GrpcClientConfig {
	return GrpcClientRawConfig(fmt.Sprintf("grpc.client.%s", name))
}

func GrpcClientRawConfig(name string) GrpcClientConfig {
	config := GrpcClientConfig{
		TraceEnable: viper.GetBool("server.grpc_trace_enable"),
		HostPort:    viper.GetString(fmt.Sprintf("%s.host_port", name)),
	}
	return config
}

func (config GrpcClientConfig) Build() (*grpc.ClientConn, io.Closer) {
	var tracer opentracing.Tracer
	var closer io.Closer
	log.Info("confgi", config)
	if config.TraceEnable {
		//开启链路追踪
		tracer, closer, _ = trace.StdConfig().Build()
		conn, _ := grpc.Dial(config.HostPort, grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
			grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)))
		return conn, closer
	}
	conn, _ := grpc.Dial(config.HostPort, grpc.WithInsecure())
	return conn, closer
}
