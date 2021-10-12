package v1

import (
	"github.com/dierbei/blind-box/internal/middleware"
	"github.com/dierbei/blind-box/pkg/wx"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func UserRegister(router *gin.RouterGroup) {
	user := UserController{}

	manGroup := router.Group("/user")
	manGroup.GET("/login", user.Login)
}

func (user *UserController) Login(ctx *gin.Context) {
	code := ctx.Query("code")
	userSessionInfo, err := wx.WxLogin(code)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	middleware.ResponseSuccess(ctx, userSessionInfo)
}
