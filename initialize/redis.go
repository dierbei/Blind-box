package initialize

import (
	"github.com/dierbei/blind-box/global"
	"github.com/go-redis/redis/v8"
)

func initRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.ServerConfig.Redis.Addr,
		Password: global.ServerConfig.Redis.Password, // no password set
		DB:       global.ServerConfig.Redis.DB,       // use default DB
	})
}
