package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/dierbei/blind-box/config"
)

var (
	// 总配置
	ServerConfig = &config.ServerConfig{}
	// MySQL连接
	MySQLTx *gorm.DB
	// Redis连接
	RedisClient *redis.Client
	// 阿里OSS
	AliOSS *oss.Client
)
