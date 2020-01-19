package view

import (
	"firstCrawler/engine"
	"firstCrawler/frontend/model"
	common "firstCrawler/model"
	"os"
	"testing"
)

func TestSearchResultViewRender(t *testing.T) {
	SearchResultView := CreateSearchResultView("./template.html")

	page := model.SearchResult{}
	page.Hits = 100
	page.Start = 20
	data := engine.Item{
		ID:   "1325499075",
		TYPE: "zhenai",
		URL:  "http://album.zhenai.com/u/1325499075",
		Payload: common.Profile{
			Name:     "心给懂的人",
			Gender:   "女",
			Area:     "阿克苏阿克苏市",
			Age:      50,
			Marriage: "离异",
			Height:   158,
			Income:   "5-8千",
			Zodiac:   "魔羯",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, data)
	}

	file, err := os.Create("./templete.test.html")
	if err != nil {
		panic(err)
	}
	err = SearchResultView.Render(file, page)
	if err != nil {
		panic(err)
	}
}
