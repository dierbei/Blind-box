package global

import (
	"github.com/dierbei/blind-box/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ServerConfig = &config.ServerConfig{}
	MySQLTx      *gorm.DB
	RedisClient  *redis.Client
)
