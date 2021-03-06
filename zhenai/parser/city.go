package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<h2 class="post-title"><a href="(https?://[\w:/.]+)[^>]*">([^<]*)</a></h2>`)
	cityURLRe = regexp.MustCompile(`<span class="pager"><a href="(https?://[\w:/.]+)[^>]*">[^<]+</a></span>`)
)

// ParseCity 城市解析器.
func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			URL:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			URL:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
