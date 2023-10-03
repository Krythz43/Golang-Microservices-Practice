package main

import (
	"test/worker"
	"time"
)

func main() {
	job := worker.Job{
		Action: PrintPayload,
		Payload: map[string]string{
			"time": time.Now().String(),
		},
	}

	defer worker.Wait()
	job.Fire()
}
