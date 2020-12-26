package worker

type Worker struct {
	workerName string
	workerCron string
	Task       func()
}

func New() Worker {
	woker := Worker{}
	return woker
}
