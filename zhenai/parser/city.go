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
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
