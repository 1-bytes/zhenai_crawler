package engine

// 抛给数据解析器的数据结构.
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// 解析器返回数据的结构.
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser 开发过程中临时使用的占位函数.
func NilParser([]byte) ParseResult {
	return ParseResult{
		Requests: nil,
		Items:    nil,
	}
}
