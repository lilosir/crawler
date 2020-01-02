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
func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChannel(in)

	for i := 0; i < e.WorkerCounter; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %d: %v\n", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
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
