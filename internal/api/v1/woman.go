package v1

import "github.com/gin-gonic/gin"

type WomanController struct {
}

func WomanRegister(router *gin.RouterGroup) {
	woman := WomanController{}

	manGroup := router.Group("/woman")
	manGroup.GET("/add", woman.AddWoman)
}

func (woman *WomanController) AddWoman(ctx *gin.Context) {

}
