package controller

import (
	"bluelell/logic"
	"bluelell/model"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//注册
func SignUpHandlerFunc(c *gin.Context) {
	//1.获取参数、参数校验
	//将请求包重的数据绑定到结构体中
	var p model.ParamSignUp                      //得先定义一个结构体装请求包中的数据
	if err := c.ShouldBindJSON(&p); err != nil { //ShouldBindJSON只能判断数据格式是否是json和字段的类型对不对
		//请求参数有误，直接返回
		zap.L().Error("sign up with invalid param") //往日志文件中打印一条Error型的日志
		//判断是否是validator.ValidationErrors类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		} //如果是这个类型，则用翻译器翻译后再返回这个错误信息
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}
	////由于ShouldBindJSON的局限性，所以需要手动对请求中的参数信息进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
	//	zap.L().Error("sign up with invalid param") //王日志文件中打印一条Error型的日志
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//}
	//if p.Password != p.RePassword {
	//	zap.L().Error("sign up with invalid param") //王日志文件中打印一条Error型的日志
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "两次密码不一致！",
	//	})
	//}
	//2.处理业务
	logic.SignUp()
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success!",
	})
}
