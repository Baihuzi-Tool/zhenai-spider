package engine

import (
	"log"
	"zhenaiSpider/simpleSoider/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetch url %s", r.Url)

		body, e := fetcher.Fetch(r.Url)
		if e != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, e)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}
