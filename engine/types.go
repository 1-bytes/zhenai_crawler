package engine

type ParserFunc func(contents []byte, url string) ParseResult

//
// Parser
// @Description: 解析器接口，为了后面支持序列化.
//
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//
// Request
// @Description: 抛给数据解析器的数据结构.
//
type Request struct {
	URL    string
	Parser Parser
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

// NilParser 开发中占位用的 Parser.
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}
func (NilParser) Serialize() (_ string, _ interface{}) {
	return "NilParser", nil
}

//
// FuncParser
// @Description: 城市列表/城市/用户 解析器封装.
//
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// NewFuncParser FuncParser 工厂函数.
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
