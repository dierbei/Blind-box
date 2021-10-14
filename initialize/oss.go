package initialize

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"

	"github.com/dierbei/blind-box/global"
)

func initOSS() {
	client, err := oss.New(global.ServerConfig.AliOSS.EndPoint, global.ServerConfig.AliOSS.AccessKey, global.ServerConfig.AliOSS.AccessKeySecret)
	if err != nil {
		zap.S().Errorw("oss.New failed", "message", err)
		return
	}
	global.AliOSS = client
}
