package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(https?://[\w:/.]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
