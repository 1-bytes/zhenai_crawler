package persist

import (
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"testing"
)

func TestSave(t *testing.T) {
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
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	// TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index("dating_profile_zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual model.Profile
	// resp.Source 是 JSON 格式的原始数据，这里将其拿到后进行了 JSON 反序列化
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
