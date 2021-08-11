package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// main 入口程序.
func main() {
	itemChan, err := persist.ItemSaver("dating_profile_zhenai")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
