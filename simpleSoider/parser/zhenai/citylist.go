package zhenai

import (
	"regexp"
	"zhenaiSpider/simpleSoider/engine"
)

const cityListRe = `<a [^href]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]*)</a>`

func ParseCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	for _, m := range submatch {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
