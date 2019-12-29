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
	go func() {
		s.WorkerChan <- r
	}()
}
