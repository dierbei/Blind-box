package router

import (
	v1 "github.com/dierbei/blind-box/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func InitManRouter(v1Group *gin.RouterGroup) {
	manGroup := v1Group.Group("/man")

	manGroup.GET("/add", v1.AddMan)
}
