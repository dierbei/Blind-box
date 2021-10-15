package initialize

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dierbei/blind-box/global"
	v1 "github.com/dierbei/blind-box/internal/api/v1"
	"github.com/dierbei/blind-box/internal/middleware"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	InitLogger()
	InitConfig()
	InitMySQL()
	initRedis()
	initOSS()

	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)

	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           global.ServerConfig.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(global.ServerConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.ServerConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(global.ServerConfig.MaxHeaderBytes),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", global.ServerConfig.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", global.ServerConfig.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}

func InitRouter() *gin.Engine {

	engine := gin.Default()

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	manGroup := engine.Group("/v1")
	manGroup.Use(
		middleware.TranslationMiddleware(),
	)
	{
		v1.ManRegister(manGroup)
	}

	userGroup := engine.Group("/v1")
	userGroup.Use(
		middleware.Logging())
	{
		v1.UserRegister(userGroup)
	}

	return engine
}
