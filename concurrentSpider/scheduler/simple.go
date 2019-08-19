package scheduler

import "zhenaiSpider/concurrentSpider/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleSchedule) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
