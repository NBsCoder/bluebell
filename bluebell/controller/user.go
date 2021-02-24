package controller

import (
	"bluelell/dao/mysql"
	"bluelell/logic"
	"bluelell/model"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandlerFunc 注册
func SignUpHandlerFunc(c *gin.Context) {
	//1.获取参数、参数校验
	//将请求包重的数据绑定到结构体中
	var p *model.ParamSignUp                    //得先定义一个结构体装请求包中的数据
	if err := c.ShouldBindJSON(p); err != nil { //ShouldBindJSON只能判断数据格式是否是json和字段的类型对不对
		//请求参数有误，直接返回
		zap.L().Error("sign up with invalid param") //往日志文件中打印一条Error型的日志
		//判断是否是validator.ValidationErrors类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		//如果是这个类型，则用翻译器翻译后再返回这个错误信息
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.处理业务
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("login.signup failed", zap.Error(err)) //这里也要往日志文件中记录日志信息
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandlerFunc 登陆
func LoginHandlerFunc(c *gin.Context) {
	//1.获取请求

	//zhelishimeiyongbufen
	//定义一个结构体获取请求包中的数据
	//将请求中的数据绑定到结构体中
	p := new(model.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回
		//往日志文件中打印一条相关日志
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断是否是validator.ValidationErrors类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.处理业务
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}
