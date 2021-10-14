package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/dierbei/blind-box/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		zap.S().Errorw("v.ReadInConfig failed", "msg", err.Error())
		return
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Errorw("v.Unmarshal failed", "msg", err.Error())
		return
	}
	zap.S().Infow("配置文件读取成功", "data", global.ServerConfig)

	//sc := []constant.ServerConfig{
	//	{
	//		IpAddr: global.ServerConfig.Host,
	//		Port:   uint64(global.ServerConfig.Port),
	//	},
	//}
	//
	//cc := constant.ClientConfig{
	//	NamespaceId:         global.NacosConfig.Namespace,
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "tmp/nacos/log",
	//	CacheDir:            "tmp/nacos/cache",
	//	RotateTime:          "1h",
	//	MaxAge:              3,
	//	LogLevel:            "debug",
	//}

	//nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": sc,
	//	"clientConfig":  cc,
	//})
	//if err != nil {
	//	zap.S().Errorw("clients.CreateConfigClient failed", "msg", err.Error())
	//	return
	//}

	//content, err := nacosClient.GetConfig(vo.ConfigParam{
	//	DataId: "goods_web.json",
	//	Group:  global.NacosConfig.Group,
	//})
	//if err != nil {
	//	zap.S().Errorw("nacosClient.GetConfig failed", "msg", err.Error())
	//	return
	//}

	//if err = json.Unmarshal([]byte(content), &global.ServerConfig); err != nil {
	//	zap.S().Errorw("json.Unmarshal failed", "msg", err.Error())
	//	return
	//}
	//zap.S().Infow("读取Nacos配置成功", "data", global.ServerConfig)

	//err = nacosClient.ListenConfig(vo.ConfigParam{
	//	DataId: "goods_web.json",
	//	Group:  global.NacosConfig.Group,
	//	OnChange: func(namespace, group, dataId, data string) {
	//		zap.S().Infow("配置文件发生变化，即将重新读取配置")
	//		if err = json.Unmarshal([]byte(content), &global.ServerConfig); err != nil {
	//			zap.S().Errorw("json.Unmarshal failed", "msg", err.Error())
	//			return
	//		}
	//		zap.S().Infow("读取Nacos配置成功", "data", global.ServerConfig)
	//	},
	//})
}
