package main

import (
	"zhenaiSpider/concurrentSpider/engine"
	"zhenaiSpider/concurrentSpider/parser/zhenai"
	"zhenaiSpider/concurrentSpider/scheduler"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleSchedule{},
		WorkerCount: 10,
	}

	request := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: zhenai.ParserCityList,
	}

	e.Run(request)
}
