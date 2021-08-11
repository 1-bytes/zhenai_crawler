package engine

import (
	"log"
)

type SimpleEngine struct{}

// Run 单任务版 engine.
func (s SimpleEngine) Run(seeds ...Request) {
	var (
		requests []Request
	)
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
