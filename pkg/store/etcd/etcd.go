package etcd

import (
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

type ConfigEtcd struct {
	Endpoints   string
	DialTimeout int64
}

//标准配置
func StdConfig() ConfigEtcd {
	return rawConfig(fmt.Sprintf("system.etcd"))
}

func rawConfig(name string) ConfigEtcd {
	config := ConfigEtcd{
		Endpoints:   viper.GetString(fmt.Sprintf("%s.dndpoints", name)),
		DialTimeout: viper.GetInt64(fmt.Sprintf("%s.dial_timeout", name)),
	}
	return config
}

func (stdConfig ConfigEtcd) Build() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(stdConfig.Endpoints, ","),
		DialTimeout: time.Duration(stdConfig.DialTimeout) * time.Second,
	})
	return cli, err
}
