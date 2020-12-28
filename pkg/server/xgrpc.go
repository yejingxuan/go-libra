package xgrpc

import (
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	Server *grpc.Server
	Listen *net.Listener
}
