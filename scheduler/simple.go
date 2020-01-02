package scheduler

import (
	"firstCrawler/engine"
)

//SimpleScheduler struct
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

//ConfigureMasterWorkerChannel add the channel into sheduler
func (s *SimpleScheduler) ConfigureMasterWorkerChannel(c chan engine.Request) {
	s.WorkerChan = c
}

//Submit requests down to work channel
func (s *SimpleScheduler) Submit(r engine.Request) {
	//important, fix loop waiting循环等待, create a go routine for every request
	// the amount of go routines is different with the worker go routine(workerCounter set in main.go)
	go func() {
		s.WorkerChan <- r
	}()
}
