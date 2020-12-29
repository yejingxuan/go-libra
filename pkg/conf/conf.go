package conf

import (
	"github.com/spf13/viper"
)

func InitConfig(path string) error {
	loadDefaultConfig()
	//加载本地配置
	viper.SetConfigName("config") // name of conf file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the conf file does not have the extension in the name
	viper.AddConfigPath(path)
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the conf file
	if err != nil {             // Handle errors reading the conf file
		println("error", err.Error())
	}
	//加载分布式配置中心 todo

	println("配置文件加载成功")
	return nil
}

func loadDefaultConfig() {
	viper.SetDefault("general.app_name", "libra")
	viper.SetDefault("general.log_level", 1)
	viper.SetDefault("general.log_path", "./log/")

	viper.SetDefault("server.http_port", 9001)
	viper.SetDefault("server.grpc_port", 8001)
	viper.SetDefault("server.grpc_protocol", "tcp")
	viper.SetDefault("server.grpc_trace_enable", false)

	viper.SetDefault("worker.name", "libra")
	viper.SetDefault("worker.corn", "0/2 * * * * ?")

	viper.SetDefault("system.trace.serviceName", "libra")
	viper.SetDefault("hostPort", "127.0.0.1:6831")
}
