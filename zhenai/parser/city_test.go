package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	// test target should be staic, because here is only testing for parse
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	contents, err := ioutil.ReadFile("./city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCity(contents, "")
	fmt.Printf("%s\n", contents)
	// verify result

	const resultSize = 20
	expectedURLs := []string{
		"http://album.zhenai.com/u/1910207788",
		"http://album.zhenai.com/u/1312667054",
		"http://album.zhenai.com/u/1477405453",
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
