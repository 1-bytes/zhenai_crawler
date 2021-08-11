package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
)

// ItemSaver 用于存储 item.
func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

// save 存储数据.
func save(item engine.Item) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	indexService := client.Index().
		Index("dating_profile_zhenai").
		BodyJson(item)
	if item.Id != "" {
		indexService = indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
