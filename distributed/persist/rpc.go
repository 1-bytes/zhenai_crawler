package persist

import (
	"crawler/engine"
	"crawler/persist"
	"github.com/olivere/elastic/v7"
	"log"
)

//
// ItemSaverService
// @Description: 用于 RPC
//
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save 用 RPC 完成 ES 存储任务
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err != nil {
		log.Printf("Error saving item %v: %v", item, err)
		return err
	}
	*result = "ok"
	return nil
}
