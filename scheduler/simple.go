package scheduler

import "crawler/engine"

//
// SimpleScheduler
// @Description: 多个 worker 公用一个 channel 来实现的调度器
//
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// WorkerChan 返回 worker 公用的 channel.
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady 实现接口.
func (s *SimpleScheduler) WorkerReady(_ chan engine.Request) {}

// Run 创建一个 worker 公用的 channel.
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
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
