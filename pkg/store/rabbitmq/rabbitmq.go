package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

type ConfigRabbitMQ struct {
	Endpoints   string
	DialTimeout int64
}

//标准配置
func StdConfig() ConfigRabbitMQ {
	return rawConfig(fmt.Sprintf("system.rabbitmq"))
}

func rawConfig(name string) ConfigRabbitMQ {
	config := ConfigRabbitMQ{
		Endpoints:   viper.GetString(fmt.Sprintf("%s.dndpoints", name)),
		DialTimeout: viper.GetInt64(fmt.Sprintf("%s.dial_timeout", name)),
	}
	return config
}

func (stdConfig ConfigRabbitMQ) Build() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(stdConfig.Endpoints, ","),
		DialTimeout: time.Duration(stdConfig.DialTimeout) * time.Second,
	})
	return cli, err
}
