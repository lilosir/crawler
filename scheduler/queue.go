package scheduler

import (
	"firstCrawler/engine"
)

//QueueScheduler struct
type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

//ConfigureMasterWorkerChannel ...
func (s *QueueScheduler) ConfigureMasterWorkerChannel(c chan engine.Request) {
}

//WorkerReady ... send request channel to worker channel
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

//Submit reques to request channel
func (s *QueueScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

//Run ..
func (s *QueueScheduler) Run() {
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			//all the go routine scenarios
			select {
			case r := <-s.RequestChan:
				// send r to a worker, but don't know if there is a worker, so put into a queue
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				// send next request to w, don't know what is next request, so put into a queue
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
