package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "原来无话可说爱你")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
