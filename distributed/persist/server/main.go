package main

import (
	configs "crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"crawler/pkg/config"
	"github.com/olivere/elastic/v7"
)

// main ItemSaver 服务端.
func main() {
	configs.Initialize()
	err := serveRPC(":"+config.GetString("app.item_saver_port"),
		config.GetString("elasticSearch.index"))
	if err != nil {
		panic(err)
	}
}

// serveRPC 启动 RPC 服务端.
func serveRPC(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
