package engine

import (
	"log"
	"zhenaiSpider/concurrentSpider/fetcher"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seed {
		if isDuplication(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
		}
		itemCount++
		for _, r := range result.Requests {
			if isDuplication(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}

}

var parsedUrl = make(map[string]bool)

func isDuplication(url string) bool {
	if exist, ok := parsedUrl[url]; ok && exist {
		return true
	}

	parsedUrl[url] = true
	return false
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetch url %s", r.Url)
	body, e := fetcher.Fetch(r.Url)

	if e != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, e)
		return ParserResult{}, nil
	}

	return r.ParserFunc(body), nil
}
