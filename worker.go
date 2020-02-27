package mypool

type Runnable interface {
	Run()
}

type Job struct {
	Fn func()
}

func (j *Job) Run() {
	j.Fn()
}

var JobQueue chan Job

type Worker struct {
	// workers
	WorkerPool chan chan Job
	// to do jobs channel
	JobChannel chan Job
	quit       chan bool
}

// new worker in this pool
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			// get job from job channel
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// run job
				job.Run()
			case <-w.quit:
				// shutdown
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
