package persist

import (
	"context"
	"errors"
	"firstCrawler/engine"
	"log"

	"github.com/olivere/elastic"
)

//ItemSaver save the items from every request
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item $%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item save: error saving item %v: %v", item, err)
			}

		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {

	if item.TYPE == "" {
		return errors.New("Must supply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.TYPE).
		BodyJson(item)

	if item.ID != "" {
		indexService = indexService.Id(item.ID)
	}

	_, error := indexService.
		Do(context.Background())
	if error != nil {
		return error
	}

	return nil
}
