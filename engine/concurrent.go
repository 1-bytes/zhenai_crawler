package engine

import (
	"crawler/fetcher"
	"crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

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

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got item#%d: %v", profileCount, item)
				profileCount++
			}
		}

		for _, request := range result.Requests {
			// URL 去重
			if isDuplicate(request.URL) {
				continue
			}

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
			result, err := e.worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// worker 任务处理程序，将 fetch 拿到的数据扔给指定的 parser 函数.
func (e *ConcurrentEngine) worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
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
