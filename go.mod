module github.com/yejingxuan/go-libra

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.3 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
