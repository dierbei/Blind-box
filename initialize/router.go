package initialize

import (
	"github.com/dierbei/blind-box/internal/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})

	v1Group := engine.Group("/v1")
	router.InitManRouter(v1Group)
	router.InitWomanRouter(v1Group)

	return engine
}
