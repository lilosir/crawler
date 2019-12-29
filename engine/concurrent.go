package engine

import "log"

//ConcurrentEngine struct
type ConcurrentEngine struct {
	Scheduler     Scheduler
	WorkerCounter int
}

//Scheduler is the interface shedule all the requests
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChannel(chan Request)
}

//Run the engine as long as there is at least one request
func (c *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChannel(in)

	for i := 0; i < c.WorkerCounter; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v\n", item)
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				log.Fatalf("wtf %s", err)
				continue
			}
			out <- result
		}
	}()
}
