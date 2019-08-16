package main

import (
	"zhenaiSpider/simpleSoider/engine"
	"zhenaiSpider/simpleSoider/parser/zhenai"
)

func main() {
	testRequest := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: zhenai.ParseCityList,
	}

	engine.Run(testRequest)
}
