package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dierbei/blind-box/global"
	"github.com/dierbei/blind-box/internal/middleware"
	"github.com/dierbei/blind-box/internal/model"
	"github.com/dierbei/blind-box/pkg/forms"
)

type ManController struct {
}

func ManRegister(router *gin.RouterGroup) {
	man := ManController{}

	manGroup := router.Group("/man")
	manGroup.POST("/add", man.AddMan)
}

func (man *ManController) AddMan(ctx *gin.Context) {
	params := &forms.ManAddForm{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	err := (&model.Man{Username: params.Username}).Insert(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, "")
	return
}
