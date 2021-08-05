package parser

import (
	"io/ioutil"
	"testing"
)

// TestParseCityList 城市列表解析器测试案例.
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s; but %s", i, url, result.Requests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("expected ciry #%d: %s; but %s", i, city, result.Items[i].(string))
		}
	}
}
