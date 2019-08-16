package zhenai

import (
	"regexp"
	"zhenaiSpider/simpleSpider/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^</]+)</a>`

func ParserCity(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	submatch := re.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, m := range submatch {
		 name := string(m[2])
		result.Items = append(result.Items, "User: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c, name)
			},
		})

	}

	return result
}
