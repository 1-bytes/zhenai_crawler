package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<h2 class="post-title"><a href="(https?://[\w:/.]+)[^>]*">([^<]*)</a></h2>`

// ParseCity 城市解析器.
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}
	return result
}
