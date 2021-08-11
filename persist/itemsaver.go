package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
)

// ItemSaver 用于存储 item.
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// save 将数据存储至 ElasticSearch.
func save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().
		Index(index).
		BodyJson(item)
	if item.ID != "" {
		indexService = indexService.Id(item.ID)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
