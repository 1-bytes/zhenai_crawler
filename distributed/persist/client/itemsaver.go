package client

import (
	"crawler/distributed/rpcsupport"
	"crawler/engine"
	"crawler/pkg/config"
	"log"
)

// ItemSaver 用于存储 item.
func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			result := ""
			err = client.Call(config.GetString("app.item_saver_rpc"), item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
