package scheduler

import (
	"crawler/engine"
)

//
// QueuedScheduler
// @Description: 各个 Worker 拥有自己的 channel 来实现的调度器
//
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(_ chan engine.Request) {
	panic("implement me")
}

// Submit 将 Request 加入到队列，等待出现空闲 Worker 的出现.
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkerReady 将空闲的 Worker 加入到队列，准备 Request 的到来.
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// Run 用来启动 Queue 版本的 Scheduler.
func (s *QueuedScheduler) Run() {
	// 初始化 channel，用于接收即将到来的 Request 和 Worker
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		// 初始化队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeReqeust engine.Request
			var activeWorker chan engine.Request
			// 从队列里拿到数据先不要删除，等成功匹配以后再去干掉该数据
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeReqeust = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case r := <-s.workerChan:
				workerQ = append(workerQ, r)
			case activeWorker <- activeReqeust:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
