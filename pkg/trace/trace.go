package trace

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

type ConfigTrace struct {
	ServiceName string
	HostPort    string
}

//标准配置
func StdConfig() ConfigTrace {
	return rawConfig(fmt.Sprintf("system.trace"))
}

func rawConfig(name string) ConfigTrace {
	config := ConfigTrace{
		ServiceName: viper.GetString(fmt.Sprintf("%s.serviceName", name)),
		HostPort:    viper.GetString(fmt.Sprintf("%s.hostPort", name)),
	}
	return config
}

func (stdConfig ConfigTrace) Build() (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: stdConfig.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: stdConfig.HostPort,
		},
	}
	//tracer, closer, err := cfg.New(stdConfig.ServiceName, config.Logger(jaeger.StdLogger))
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return tracer, closer, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}
