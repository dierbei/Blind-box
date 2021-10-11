package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/dierbei/blind-box/global"
)

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ServerConfig.MySQL.User,
		global.ServerConfig.MySQL.Password,
		global.ServerConfig.MySQL.Host,
		global.ServerConfig.MySQL.Port,
		global.ServerConfig.MySQL.DbName,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	var err error
	global.MySQLTx, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Errorw("gorm.Open failed. err:", "msg", err.Error())
	}
}
