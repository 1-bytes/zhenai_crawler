package scheduler

import (
	"crawler/engine"
)

//
// QueuedScheduler
// @Description: 各个 worker 拥有自己的 channel 来实现的调度器
//
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// WorkerChan 为每一个 worker 创建一个 channel
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// Submit 将 request 加入到队列，等待出现空闲 worker 的出现.
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkerReady 将空闲的 worker 加入到队列，准备 request 的到来.
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// Run 用来启动 queue 版本的 scheduler.
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
