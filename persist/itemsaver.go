package persist

import (
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
)

// ItemSaver 用于存储 item.
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

// save 存储数据.
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("dating_profile_zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}