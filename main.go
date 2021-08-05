package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

// main ...
func main() {
	engine.Run(engine.Request{
		URL:        "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
