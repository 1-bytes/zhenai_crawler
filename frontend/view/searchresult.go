package view

import (
	"crawler/frontend/model"
	"html/template"
	"io"
)

//
// SearchResultView
// @Description: 模板渲染.
//
type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView 初始化模板并返回.
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

// Render 模板渲染.
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
