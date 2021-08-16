package persist

import (
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"testing"
)

// TestSave 测试数据存储.
func TestSave(t *testing.T) {
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

	// TODO: Try to start up elastic search here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Id(expected.ID).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	// resp.Source 是 JSON 格式的原始数据，这里将其拿到后进行了 JSON 反序列化
	_ = json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
