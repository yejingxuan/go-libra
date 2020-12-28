package worker

import (
	"github.com/spf13/viper"
)

type Worker struct {
	Config WorkerConfig
	task   func()
}

type WorkerConfig struct {
	WorkerName string
	WorkerCron string
}

//默认配置
func DefaultConfig() WorkerConfig {
	config := WorkerConfig{
		WorkerName: viper.GetString("worker.name"),
		WorkerCron: viper.GetString("worker.corn"),
	}
	return config
}

//标准配置
func StdConfig(name string) WorkerConfig {
	config := WorkerConfig{
		WorkerName: viper.GetString("worker." + name + ".name"),
		WorkerCron: viper.GetString("worker." + name + ".corn"),
	}
	return config
}

//自定义配置
func RawConfig(name string) WorkerConfig {
	config := WorkerConfig{
		WorkerName: viper.GetString(name + ".name"),
		WorkerCron: viper.GetString(name + ".corn"),
	}
	return config
}

func (config WorkerConfig) Build(task func()) Worker {
	worker := Worker{
		Config: config,
		task:   task,
	}
	return worker
}

func (worker Worker) GetTask() func() {
	return worker.task
}
