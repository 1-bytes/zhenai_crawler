package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(https?://[\w:/.]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市列表解析器.
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			URL:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
