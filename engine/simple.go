package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

// Run 爬虫核心调度引擎.
func (s SimpleEngine) Run(seeds ...Request) {
	var (
		requests []Request
	)
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := s.worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// worker .
func (s SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
