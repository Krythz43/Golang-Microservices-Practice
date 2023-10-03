package main

import (
	"test/worker"
	"time"
)

func main() {
	defer worker.Wait()

	for i := 0; i < 100; i++ {
		job := worker.Job{
			Action: PrintPayload,
			Payload: map[string]string{
				"time": time.Now().String(),
			},
		}

		// This generic job here can be serialized and deserialized using gob package
		// https://stackoverflow.com/questions/28020070/golang-serialize-and-deserialize-back/30721381#30721381
		job.Fire()
	}
}
