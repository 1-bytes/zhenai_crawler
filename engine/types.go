package engine

// Request 抛给数据解析器的数据结构.
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult 解析器返回数据的结构.
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	URL     string
	Id      string
	Payload interface{}
}

// NilParser 开发过程中临时使用的占位函数.
// func NilParser(_ []byte) ParseResult {
// 	return ParseResult{
// 		Requests: nil,
// 		Items:    nil,
// 	}
// }
