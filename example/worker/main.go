package main

import (
	libra "github.com/yejingxuan/go-libra/pkg"
	"github.com/yejingxuan/go-libra/pkg/log"
	"github.com/yejingxuan/go-libra/pkg/worker"
)

func main() {
	app := libra.DefaultApplication()
	app.Start()
	app.AppendWorkes(weatherWorker(), eatWorker())
	app.Run()
}

//天气预报任务
func weatherWorker() worker.Worker {
	worker := worker.StdConfig("weather").Build(func() {
		log.Info("任务开始执行111")
	})
	return worker
}

//团建任务
func eatWorker() worker.Worker {
	worker := worker.StdConfig("eat").Build(func() {
		log.Info("任务开始执行222")
	})
	return worker
}
