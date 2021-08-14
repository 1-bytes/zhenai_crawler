package model

import "crawler/engine"

//
// SearchResult
// @Description: 搜索和分页使用的结构体.
//
type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}
