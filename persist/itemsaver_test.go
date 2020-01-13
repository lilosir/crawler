package persist

import (
	"context"
	"encoding/json"
	"firstCrawler/engine"
	"firstCrawler/model"
	"fmt"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		ID:   "113",
		TYPE: "zhenai",
		URL:  "http://1231",
		Payload: model.Profile{
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

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	// save expected item
	const index = "dating_test"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//get saved item
	result, err := client.Get().
		Index(index).
		Type(expected.TYPE).
		Id(expected.ID).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", result.Source)
	var actual engine.Item

	err = json.Unmarshal(*result.Source, &actual)

	parsedProfile, _ := model.FromJSONObj(actual.Payload)
	actual.Payload = parsedProfile

	//verify result
	if expected != actual {
		t.Errorf("got %v, but expect %v", actual, expected)
	}
}
