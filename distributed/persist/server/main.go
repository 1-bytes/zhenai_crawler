package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"github.com/olivere/elastic/v7"
)

func main() {
	err := serveRPC(config.ItemSaverPort, config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

// serveRPC 启动 RPC 服务端
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
