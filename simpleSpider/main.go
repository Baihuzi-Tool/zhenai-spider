package main

import (
	"zhenaiSpider/simpleSpider/engine"
	"zhenaiSpider/simpleSpider/parser/zhenai"
)

func main() {
	testRequest := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: zhenai.ParserCityList,
	}

	engine.Run(testRequest)
}
