package engine

import (
	"firstCrawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body, r.URL), nil
}
