package zhenai

import (
	"regexp"
	"zhenaiSpider/simpleSoider/engine"
)

const cityListRe = `<a [^href]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]*)</a>`

func ParserCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(content, -1)

	limit := 1
	result := engine.ParserResult{}
	for _, m := range submatch {
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})

		limit--
		if (limit <= 0) {
			break
		}
	}

	return result
}
