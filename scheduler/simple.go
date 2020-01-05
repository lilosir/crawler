package scheduler

import (
	"firstCrawler/engine"
)

//SimpleScheduler struct
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

//Submit requests down to work channel
func (s *SimpleScheduler) Submit(r engine.Request) {
	//important, fix loop waiting循环等待, create a go routine for every request
	// the amount of go routines is different with the worker go routine(workerCounter set in main.go)
	go func() {
		s.WorkerChan <- r
	}()
}

//Run creates the only worker channel in simple engine
func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

// CreateWorkerChan will return worker channel, simple schedule only has one
func (s *SimpleScheduler) CreateWorkerChan() chan engine.Request {
	return s.WorkerChan
}

//WorkerReady doing nothing in simple sheduler
func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}
