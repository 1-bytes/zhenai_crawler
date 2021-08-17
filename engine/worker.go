package engine

import (
	"crawler/fetcher"
	"log"
)

// worker .
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.URL), nil
}
