package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// main 入口程序.
func main() {
	itemChan, err := client.ItemSaver(config.ItemSaverPort)
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
