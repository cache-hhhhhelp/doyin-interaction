package config

import (
		"github.com/zeromicro/go-zero/zrpc"
		"github.com/zeromicro/go-zero/core/stores/cache"
)
type Config struct {
	zrpc.RpcServerConf
	Cache cache.CacheConf
	Mysql struct {
		Datasource string
	}
}
