package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	common "crawler/model"
	"os"
	"testing"
)

// TestSearchResultView_Render 测试模板渲染.
func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		URL: "http://localhost:8080/mock/album.zhenai.com/u/803409738748213657",
		ID:  "803409738748213657",
		Payload: common.Profile{
			Name:       "原来无话可说爱你",
			Gender:     "女",
			Age:        82,
			Height:     82,
			Weight:     173,
			Income:     "10001-20000元",
			Marriage:   "离异",
			Education:  "高中",
			Occupation: "人事/行政",
			Hokou:      "青岛市",
			Xingzuo:    "天蝎座",
			House:      "租房",
			Car:        "有豪车",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}
}
