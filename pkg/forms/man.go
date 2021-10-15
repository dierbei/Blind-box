package forms

import (
	"github.com/gin-gonic/gin"

	"github.com/dierbei/blind-box/pkg/public"
)

type ManListPageInput struct {
	PageSize int `form:"page_size" json:"page_size" comment:"每页记录数" validate:"" example:"10"`
	Page     int `form:"page" json:"page" comment:"页数" validate:"required" example:"1"`
}

func (params *ManListPageInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type PeopleAddForm struct {
	WxNumber    string `form:"wx_number" json:"wx_number" comment:"微信号" validate:"required"`
	Description string `form:"description" json:"description"  comment:"个人简介" validate:"required"`
	Local       string `form:"local" json:"local" comment:"位置" validate:"required"`
	//Images      string `form:"fileList" json:"fileList" comment:"自拍"`
}

func (params *PeopleAddForm) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
