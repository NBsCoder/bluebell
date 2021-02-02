package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 返回响应结构体
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseError 返回错误信息
func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMsg 返回自定义的一些错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 返回统一的成功的信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
