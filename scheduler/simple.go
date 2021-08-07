package scheduler

import "crawler/engine"

//
// SimpleScheduler
// @Description: 多个 worker 公用一个 channel 来实现的调度器
//
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan 配置任务调度器.
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit 用于将任务交至任务调度器.
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
