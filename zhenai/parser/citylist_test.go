package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// test target should be staic, because here is only testing for parse
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("./citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")
	// verify result

	// const resultSize = 470
	const resultSize = 5
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}

	for index, url := range expectedURLs {
		if url != result.Requests[index].URL {
			t.Errorf("expected url #%d: %s; but got %s\n", index, url, result.Requests[index].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Items))
	}
}
