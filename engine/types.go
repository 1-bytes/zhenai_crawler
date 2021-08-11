package engine

type ParserFunc func(contents []byte, url string) ParseResult

//
// Request
// @Description: 抛给数据解析器的数据结构.
//
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

//
// ParseResult
// @Description: 解析器返回数据的结构.
//
type ParseResult struct {
	Requests []Request
	Items    []Item
}

//
// Item
// @Description: 通用模板 后续要存储数据的格式结构.
//
type Item struct {
	URL     string
	ID      string
	Payload interface{}
}
