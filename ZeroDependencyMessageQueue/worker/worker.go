package worker

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Job struct {
	Id1     string
	Action  func(map[string]string)
	Payload map[string]string
}

func Wait() {
	wg.Wait()
}

func (job Job) Fire() {
	fmt.Println("In fire function")
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		job.Action(job.Payload)
		wg.Done()
	}()
}
