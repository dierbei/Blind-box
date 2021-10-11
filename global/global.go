package global

import (
	"github.com/dierbei/blind-box/config"
	"gorm.io/gorm"
)

var (
	ServerConfig = &config.ServerConfig{}
	MySQLTx      *gorm.DB
)
