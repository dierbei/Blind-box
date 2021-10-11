package forms

import (
	"github.com/dierbei/blind-box/pkg/public"
	"github.com/gin-gonic/gin"
)

type ManListPageInput struct {
	PageSize int `form:"page_size" json:"page_size" comment:"每页记录数" validate:"" example:"10"`
	Page     int `form:"page" json:"page" comment:"页数" validate:"required" example:"1"`
}

func (params *ManListPageInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ManAddForm struct {
	Username string `form:"username" json:"username" comment:"用户名" validate:"required"`
}

func (params *ManAddForm) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
