package main

import (
	configs "crawler/distributed/config"
	"crawler/engine"
	"crawler/persist"
	"crawler/pkg/config"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// main 入口程序.
func main() {
	configs.Initialize()
	itemChan, err := persist.ItemSaver(config.GetString("elasticSearch.index"))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		URL:    "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
