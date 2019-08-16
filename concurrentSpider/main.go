package main

import (
	"zhenaiSpider/concurrentSpider/engine"
	"zhenaiSpider/concurrentSpider/parser/zhenai"
)

func main() {
	request := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: zhenai.ParserCityList,
	}

	engine.Run(request)
}
