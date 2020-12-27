package worker

type Worker struct {
	Config WorkerConfig
	task   func()
}

type WorkerConfig struct {
	WorkerName string
	WorkerCron string
}

func StdConfig() WorkerConfig {
	config := WorkerConfig{
		WorkerName: "woker-libra",
		WorkerCron: "0/2 * * * * ?",
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
