package engine

type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{
		Requests: nil,
		Items:    nil,
	}
}
