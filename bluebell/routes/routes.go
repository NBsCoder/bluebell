package routes

import (
	"bluelell/controller"
	"bluelell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//用上两个中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	//注册
	r.POST("/signup", controller.SignUpHandlerFunc)
	//登陆
	r.POST("/login", controller.LoginHandlerFunc)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
