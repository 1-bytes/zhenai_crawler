package engine

import (
	"crawler/fetcher"
	"log"
)

// Run 爬虫核心调度引擎.
func Run(seeds ...Request) {
	var (
		requests []Request
	)
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
