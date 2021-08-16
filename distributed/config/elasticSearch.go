package config

import "crawler/pkg/config"

// init 初始化 ElasticSearch 配置.
func init() {
	config.Add("elasticSearch", config.StrMap{
		"index": config.Env("ELASTICSEARCH_INDEX", "dating_profile_zhenai"),
	})
}
