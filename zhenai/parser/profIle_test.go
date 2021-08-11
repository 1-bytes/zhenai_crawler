package parser

import (
	"crawler/engine"
	"crawler/model"
	"io/ioutil"
	"testing"
)

// TestParseProfile 测试用户信息解析器.
func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,
		"http://localhost:8080/mock/album.zhenai.com/u/803409738748213657",
		"原来无话可说爱你")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}
	actual := result.Items[0]
	expected := engine.Item{
		URL: "http://localhost:8080/mock/album.zhenai.com/u/803409738748213657",
		ID:  "803409738748213657",
		Payload: model.Profile{
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

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
