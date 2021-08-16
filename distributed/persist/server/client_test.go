package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"crawler/pkg/config"
	"testing"
)

// TestItemSaver 测试 ItemSaver RPC.
func TestItemSaver(t *testing.T) {
	host := ":" + config.GetString("app.item_saver_port")

	go func(host string) {
		err := serveRPC(host,
			"test_"+config.GetString("elasticSearch.index"))
		if err != nil {
			panic(err)
		}
	}(host)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
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
	result := ""
	err = client.Call(config.GetString("app.item_saver_rpc"), item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
