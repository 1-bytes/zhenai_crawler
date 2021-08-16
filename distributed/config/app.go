package config

import "crawler/pkg/config"

// init 初始化应用基础配置.
func init() {
	config.Add("app", config.StrMap{
		"debug":           config.Env("APP_DEBUG", false),
		"item_saver_port": config.Env("APP_ITEM_SAVER_PORT", 1234),
		"item_saver_rpc":  config.Env("APP_ITEM_SAVER_RPC", "ItemSaverService.Save"),
	})
}
