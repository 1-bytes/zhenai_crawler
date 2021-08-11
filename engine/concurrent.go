package engine

//
// ConcurrentEngine
// @Description: 并发版 Engine 基础管理套件.
//
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

//
// Scheduler
// @Description: 任务调度接口.
//
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

//
// ReadyNotifier
// @Description: work 任务完成通知接口.
//
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run 并发版 engine.
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		// 数据存储
		for _, item := range result.Items {
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		for _, request := range result.Requests {
			// URL 去重
			if isDuplicate(request.URL) {
				continue
			}

			// 提交至 request 队列，准备进入 worker 进一步处理
			e.Scheduler.Submit(request)
		}
	}
}

// createWorker 创建一个 worker.
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
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

var visitedURLs = make(map[string]bool)

// isDuplicate 简单去重，判断 url 是否重复 重复返回 true，不重复返回 false.
func isDuplicate(url string) bool {
	if visitedURLs[url] {
		return true
	}
	visitedURLs[url] = true
	return false
}
