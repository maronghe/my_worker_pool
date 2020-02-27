package main

import (
	"test/mypool"
	"time"
)

var (
	MaxWorker = 3
	MaxQueue  = 100
	JobNumber = 10
)

// 1. create a job
// 2. send to jobQueue
// 3. get job from jobQueue to JobChannel, will block to the
// 4. get data from JobChannel to WorkChannel
// 5. get data from workChannel to execute

// create a pool with maxWorker number workers
//

func main() {
	dispatcher := mypool.NewDispatcher(MaxWorker, MaxQueue)
	dispatcher.Run()

	go func() {
		for i := 0; i < JobNumber; i++ {
			job := dispatcher.MakeJob(i)
			time.Sleep(500 * time.Millisecond)
			//fmt.Println("put job")
			mypool.JobQueue <- job
		}
	}()

	time.Sleep(120 * time.Second)
}
