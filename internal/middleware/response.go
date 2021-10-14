package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
//const (
//	SuccessCode ResponseCode = iota
//	UndefErrorCode
//	ValidErrorCode
//	InternalErrorCode
//
//	InvalidRequestErrorCode ResponseCode = 401
//	CustomizeCode           ResponseCode = 1000
//
//	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
//)

type Response struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	resp := &Response{Code: code, Message: err.Error(), Data: ""}
	c.JSON(http.StatusOK, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Message: "", Data: data}
	c.JSON(http.StatusOK, resp)
}
