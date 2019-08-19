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
			e.Scheduler.Submit(r)
		}

	}

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
