package parser

import (
	"firstCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	// test target should be staic, because here is only testing for parse
	// contents, err := fetcher.Fetch("http://album.zhenai.com/u/1325499075")
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "心给懂的人")

	if len(result.Items) != 1 {
		t.Errorf("Item should only has 1 element, but got %d\n", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expectedProfile := model.Profile{
		Name:     "心给懂的人",
		Gender:   "女",
		Area:     "阿克苏阿克苏市",
		Age:      50,
		Marriage: "离异",
		Height:   158,
		Income:   "5-8千",
		Zodiac:   "魔羯",
	}

	if profile != expectedProfile {
		t.Errorf("expected %v, but got %v\n", expectedProfile, profile)
	}

}
