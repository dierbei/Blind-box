package main

import (
	"github.com/dierbei/blind-box/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

//func main() {
//	key := global.RedisClient.RandomKey(context.Background())
//	fmt.Println(key)
//}
//

func main() {
	dsn := "root:root@tcp(192.168.244.137:3306)/blindbox?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(
		//&model.Man{},
		//&model.Woman{},
		&model.People{},
		&model.Image{},
		&model.User{},
		&model.UserPrize{},
	)
}

//
//import (
//	"fmt"
//	"os"
//	"strings"
//
//	"github.com/aliyun/aliyun-oss-go-sdk/oss"
//)
//

//
//func main() {
//	// 创建OSSClient实例。
//	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "LTAI4Fyc2WMmxFpWigqrY1kM", "3lbGsC3RvOOb812qKgoeatVhMzmzYL")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//
//	// 获取存储空间。
//	bucket, err := client.Bucket("manhe-hcyj")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//
//	// 指定存储类型为标准存储，缺省也为标准存储。
//	storageType := oss.ObjectStorageClass(oss.StorageStandard)
//
//	// 指定存储类型为归档存储。
//	// storageType := oss.ObjectStorageClass(oss.StorageArchive)
//
//	// 指定访问权限为公共读，缺省为继承bucket的权限。
//	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
//
//	// 上传字符串。
//	err = bucket.PutObject("zipai.jpg", strings.NewReader(base64code), storageType, objectAcl)
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//}
