package main

import (
	"bluelell/controller"
	"bluelell/dao/mysql"
	"bluelell/dao/redis"
	"bluelell/logger"
	"bluelell/pkg/snowflake"
	"bluelell/routes"
	"bluelell/settings"
	"fmt"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: bluebell config.yaml")
	//	return
	//}
	//1、初始化配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}
	//2、初始化日志（将日志文件读到viper中）
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	//把缓存里的日志文件加载到日志文件中
	//defer zap.L().Sync()
	//zap.L().Debug("log init success!")
	//3、初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4、初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
	}
	defer redis.Close()
	//***初始化雪花算法生成用户id***
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	//***初始化gin框架内置的检验器使用的翻译器***
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed!err:%v\n", err)
		return
	}
	//5、注册路由
	r := routes.SetupRouter()
	//6、启动服务（优雅关机）
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("run server failed!err:%v\n", err)
	}

	//	go func() {
	//		// 开启一个goroutine启动服务
	//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//			log.Fatalf("listen: %s\n", err)
	//		}
	//	}()
	//
	//	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	//	// kill 默认会发送 syscall.SIGTERM 信号
	//	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	//	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	//	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	//	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	//	log.Println("Shutdown Server ...")
	//	// 创建一个5秒超时的context
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//	if err := srv.Shutdown(ctx); err != nil {
	//		log.Fatal("Server Shutdown: ", err)
	//	}
	//
	//	log.Println("Server exiting")
}

//// SignIn 用户登陆 签发token
//func SignIn(c *gin.Context) {
//	var postData SignUpReq
//	if err := c.ShouldBind(&postData); err != nil {
//		response.ReturnJSON(c, http.StatusOK, statuscode.Faillure.Code, statuscode.Faillure.Msg, err)
//		return
//	}
//	var payLoad = customerjwt.CustomClaims{TimeStr: time.Now().Format("2006-01-02 15:04:05"), Name: postData.Name, Password: postData.Pwd}
//	token, err := customerjwt.NewJWT().CreateToken(payLoad)
//	if err != nil {
//		response.ReturnJSON(c, http.StatusOK, statuscode.Faillure.Code, statuscode.Faillure.Msg, err)
//		return
//	}
//	response.ReturnJSON(c, http.StatusOK, statuscode.Faillure.Code, statuscode.Faillure.Msg, map[string]interface{}{
//		"token": token,
//	})
//}
