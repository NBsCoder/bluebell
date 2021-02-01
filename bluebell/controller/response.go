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
	rd := ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
