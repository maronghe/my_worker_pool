package mypool

import (
	"fmt"
	"math/rand"
	"time"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	// means that will create MaxWorkers number workers.
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers, maxQueue int) *Dispatcher {
	if maxQueue < 0 || maxWorkers < 0 {
		//TODO panic("can not less than zero.")
	}
	JobQueue = make(chan Job, maxQueue)     // init JobQueue
	pool := make(chan chan Job, maxWorkers) // maxWorkers's no buffered channel
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	fmt.Printf("starting %d number of workers\n", d.MaxWorkers)
	// create workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	fmt.Printf("start finished.\n")

	// dispatch Job
	go d.dispatch()
	//go d.dispatch0()
}

// base on worker to select a job
func (d *Dispatcher) dispatch() {
	for {
		select {
		case jobChannel := <-d.WorkerPool:
			//fmt.Println("get the job from work pool jobChannel ", jobChannel)
			go func(jobChannel chan Job) {
				job := <-JobQueue
				jobChannel <- job
			}(jobChannel)
		}
	}
}

// base on job to select a worker
func (d *Dispatcher) dispatch0() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func (d *Dispatcher) MakeJob(jobID int) Job {
	f := func() {
		fmt.Println("job jobId = ", jobID, " is working")
		t := randSleep(10) // t <= 10
		time.Sleep(time.Duration(t) * time.Second)
		fmt.Println("job jobId = ", jobID, " work done, keep ", t, " second.")
	}
	return Job{Fn: f}
}

func randSleep(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n) + 1
}
