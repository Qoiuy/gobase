package handler

import (
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/flog"
	"github.com/Qoiuy/go-base/backend/redis"
)

func InitRedisCache() *redis.RedisServer {
	// init redis
	redisConn := redis.NewDefaultServer(config.Opts.RedisAddress, config.Opts.RedisPassword, config.Opts.RedisDB)
	if err := redis.Check(redisConn); err != nil {
		flog.Fatalf("redis 启动参数配置错误 :%+v", err)
	}
	return redisConn
}
