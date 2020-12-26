package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron"
)

func main() {

	wait := sync.WaitGroup{}
	wait.Add(1)

	c := cron.New()
	c.AddFunc("0/2 * * * * ?", func() {
		fmt.Println("task")
	})
	c.Start()

	wait.Wait()
}
