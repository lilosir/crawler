package engine

import "log"

//ConcurrentEngine struct
type ConcurrentEngine struct {
	Scheduler     Scheduler
	WorkerCounter int
	ItemChan      chan interface{}
}

//Scheduler is the interface shedule all the requests
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	CreateWorkerChan() chan Request
	Run()
}

//ReadyNotifier interface
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//Run the engine as long as there is at least one request
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCounter; i++ {
		createWorker(e.Scheduler.CreateWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicated(r.URL) {
			log.Printf("Duplicate request: %s\n", r.URL)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(i interface{}) {
				e.ItemChan <- i
			}(item)
			// warnings below
			// go func() {
			// 	e.ItemChan <- item
			// }()
		}

		//URL dedup
		for _, request := range result.Requests {
			if isDuplicated(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedURL = make(map[string]bool)

func isDuplicated(url string) bool {
	if visitedURL[url] {
		return true
	}
	visitedURL[url] = true
	return false
}
