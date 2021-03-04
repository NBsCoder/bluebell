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

//type student struct {
//	name string
//	age  int
//}
//
//func demo() {
//	m := make(map[string]*student)
//	stus := []student{
//		{name: "小王子", age: 18},
//		{name: "小娜扎", age: 23},
//		{name: "大王八", age: 9000},
//	}
//	for _, stu := range stus {
//		m[stu.name] = &stu
//	}
//	for k, v := range m {
//		fmt.Println(k, "=>", v.name)
//	}
//}
//package main
//
//import "awesomeProject/model"
//
//func main() {
//	user:=model.NewUser("xiaoming","123456",20)
//	user.GetMoney("xiaoming","123456")
//	user.SetPwd("xiaoming","123456","111111")
//}
//package model
//
//import "fmt"
//
//type user struct {
//	name string
//	pwd string
//	money int
//}
//
//func NewUser(name string,pwd string,money int)*user{
//	//处理细节
//	if name==""{
//		fmt.Println("名字有误")
//	}
//	if pwd!="123456"{
//		fmt.Println("密码错误")
//	}
//	if money<0&&money>1000{
//		fmt.Println("穷逼")
//	}
//	return &user{
//		name:name,
//		pwd:pwd,
//		money: money,
//	}
//}
//func (u *user)SetPwd(name string,pwd string,newPwd string){
//	if name!=u.name{
//		fmt.Println("用户名输入错误")
//		return
//	}
//	if pwd!=u.pwd{
//		fmt.Println("密码输入错误")
//		return
//	}
//	u.pwd=newPwd
//	fmt.Printf("修改密码成功,新密码为：%v\n",u.pwd)
//}
//func (u user)GetMoney(name string,pwd string){
//	if name==""{
//		fmt.Println("无此用户")
//		return
//	}
//	if pwd!=u.pwd{
//		fmt.Println("密码输入错误")
//		return
//	}
//	fmt.Printf("余额：%d\n",u.money)
//}
//type A struct {
//	name string
//}
//type B struct {
//	name string
//}
//type C struct {
//	name string
//}
//type D struct {
//	A
//	B
//	c    C
//	name string
//}
//
//func Animal() {
//	var d D
//	d.name = "abc"
//	d.B.name = "def"
//	d.c.name = "ghi"
//}
//package main
//import (
//"fmt"
//)
//
////声明/定义一个接口
//type Usb interface {
//	//声明了两个没有实现的方法
//	Start()
//	Stop()
//}
//
//type Phone struct {
//	name string
//}
//
////让Phone 实现 Usb接口的方法
//func (p Phone) Start() {
//	fmt.Println("手机开始工作。。。")
//}
//func (p Phone) Stop() {
//	fmt.Println("手机停止工作。。。")
//}
//
//func (p Phone) Call() {
//	fmt.Println("手机 在打电话..")
//}
//
//type Camera struct {
//	name string
//}
////让Camera 实现   Usb接口的方法
//func (c Camera) Start() {
//	fmt.Println("相机开始工作。。。")
//}
//func (c Camera) Stop() {
//	fmt.Println("相机停止工作。。。")
//}
//
//
//func main() {
//	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
//	//这里就体现出多态数组
//	var usbArr [3]Usb
//	usbArr[0] = Phone{"vivo"}
//	usbArr[1] = Phone{"小米"}
//	usbArr[2] = Camera{"尼康"}
//
//	fmt.Println(usbArr)
//}
//package main
//
//import (
//"fmt"
//)
////定义一个结构体Account
//type Account struct {
//	AccountNo string
//	Pwd string
//	Balance float64
//}
//
////方法
////1. 存款
//func (account *Account) Deposite(money float64, pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.Pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	//看看存款金额是否正确
//	if money <= 0 {
//		fmt.Println("你输入的金额不正确")
//		return
//	}
//
//	account.Balance += money
//	fmt.Println("存款成功~~")
//
//}
//
////取款
//func (account *Account) WithDraw(money float64, pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.Pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	//看看取款金额是否正确
//	if money <= 0  || money > account.Balance {
//		fmt.Println("你输入的金额不正确")
//		return
//	}
//
//	account.Balance -= money
//	fmt.Println("取款成功~~")
//
//}
//
////查询余额
//func (account *Account) Query(pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.Pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)
//
//}
//
//
//func main() {
//
//	//测试一把
//	account := Account{
//		AccountNo : "gs1111111",
//		Pwd : "666666",
//		Balance : 100.0,
//	}
//
//	//这里可以做的更加灵活，就是让用户通过控制台来输入命令...
//	//菜单....
//	account.Query("666666")
//	account.Deposite(200.0, "666666")
//	account.Query("666666")
//	account.WithDraw(150.0, "666666")
//	account.Query("666666")
//}
//package model
//import "fmt"
//
//type person struct {
//	Name string
//	age int   //其它包不能直接访问..
//	sal float64
//}
//
////写一个工厂模式的函数，相当于构造函数
//func NewPerson(name string) *person {
//	return &person{
//		Name : name,
//	}
//}
//
////为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
//func (p *person) SetAge(age int) {
//	if age >0 && age <150 {
//		p.age = age
//	} else {
//		fmt.Println("年龄范围不正确..")
//		//给程序员给一个默认值
//	}
//}
//
//func (p *person) GetAge() int {
//	return p.age
//}
//
//
//func (p *person) SetSal(sal float64) {
//	if sal >= 3000 && sal <= 30000 {
//		p.sal = sal
//	} else {
//		fmt.Println("薪水范围不正确..")
//
//	}
//}
//
//func (p *person) GetSal() float64 {
//	return p.sal
//}
//package model
//
//import (
//"fmt"
//)
////定义一个结构体account
//type account struct {
//	accountNo string
//	pwd string
//	balance float64
//}
//
////工厂模式的函数-构造函数
//func NewAccount(accountNo string, pwd string, balance float64) *account {
//
//	if len(accountNo) < 6 || len(accountNo) > 10 {
//		fmt.Println("账号的长度不对...")
//		return nil
//	}
//
//	if len(pwd) != 6 {
//		fmt.Println("密码的长度不对...")
//		return nil
//	}
//
//	if balance < 20 {
//		fmt.Println("余额数目不对...")
//		return nil
//	}
//
//	return &account{
//		accountNo : accountNo,
//		pwd : pwd,
//		balance : balance,
//	}
//
//}
//
////方法
////1. 存款
//func (account *account) Deposite(money float64, pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	//看看存款金额是否正确
//	if money <= 0 {
//		fmt.Println("你输入的金额不正确")
//		return
//	}
//
//	account.balance += money
//	fmt.Println("存款成功~~")
//
//}
//
////取款
//func (account *account) WithDraw(money float64, pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	//看看取款金额是否正确
//	if money <= 0  || money > account.balance {
//		fmt.Println("你输入的金额不正确")
//		return
//	}
//
//	account.balance -= money
//	fmt.Println("取款成功~~")
//
//}
//
////查询余额
//func (account *account) Query(pwd string)  {
//
//	//看下输入的密码是否正确
//	if pwd != account.pwd {
//		fmt.Println("你输入的密码不正确")
//		return
//	}
//
//	fmt.Printf("你的账号为=%v 余额=%v \n", account.accountNo, account.balance)
//
//}
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	//声明一个变量，保存接收用户输入的选项
//	key := ""
//	//声明一个变量，控制是否退出for
//	loop := true
//
//	//定义账户的余额 []
//	balance := 10000.0
//	//每次收支的金额
//	money := 0.0
//	//每次收支的说明
//	note := ""
//	//收支的详情使用字符串来记录
//	//当有收支时，只需要对details 进行拼接处理即可
//	details := "收支\t账户金额\t收支金额\t说    明"
//
//	//显示这个主菜单
//	for {
//		fmt.Println("\n-----------------家庭收支记账软件-----------------")
//		fmt.Println("                  1 收支明细")
//		fmt.Println("                  2 登记收入")
//		fmt.Println("                  3 登记支出")
//		fmt.Println("                  4 退出软件")
//		fmt.Print("请选择(1-4)：")
//		fmt.Scanln(&key)
//
//		switch key {
//		case "1":
//			fmt.Println("-----------------当前收支明细记录-----------------")
//			fmt.Println(details)
//		case "2":
//			fmt.Println("本次收入金额:")
//			fmt.Scanln(&money)
//			balance += money // 修改账户余额
//			fmt.Println("本次收入说明:")
//			fmt.Scanln(&note)
//			//将这个收入情况，拼接到details变量
//			//收入    11000           1000            有人发红包
//			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
//		case "3":
//			fmt.Println("登记支出:")
//		case "4":
//			loop = false
//		default:
//			fmt.Println("请输入正确的选项...")
//		}
//
//		if !loop {
//			break
//		}
//	}
//	fmt.Println("你退出家庭记账软件的使用...")
//}
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	//声明一个变量，保存接收用户输入的选项
//	key := ""
//	//声明一个变量，控制是否退出for
//	loop := true
//
//	//定义账户的余额 []
//	balance := 10000.0
//	//每次收支的金额
//	money := 0.0
//	//每次收支的说明
//	note := ""
//	//收支的详情使用字符串来记录
//	//当有收支时，只需要对details 进行拼接处理即可
//	details := "收支\t账户金额\t收支金额\t说    明"
//
//	//显示这个主菜单
//	for {
//		fmt.Println("\n-----------------家庭收支记账软件-----------------")
//		fmt.Println("                  1 收支明细")
//		fmt.Println("                  2 登记收入")
//		fmt.Println("                  3 登记支出")
//		fmt.Println("                  4 退出软件")
//		fmt.Print("请选择(1-4)：")
//		fmt.Scanln(&key)
//
//		switch key {
//		case "1":
//			fmt.Println("-----------------当前收支明细记录-----------------")
//			fmt.Println(details)
//		case "2":
//			fmt.Println("本次收入金额:")
//			fmt.Scanln(&money)
//			balance += money // 修改账户余额
//			fmt.Println("本次收入说明:")
//			fmt.Scanln(&note)
//			//将这个收入情况，拼接到details变量
//			//收入    11000           1000            有人发红包
//			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
//		case "3":
//			fmt.Println("本次支出金额:")
//			fmt.Scanln(&money)
//			//这里需要做一个必要的判断
//			if money > balance {
//				fmt.Println("余额的金额不足")
//				break
//			}
//			balance -= money
//			fmt.Println("本次支出说明:")
//			fmt.Scanln(&note)
//			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
//		case "4":
//			loop = false
//		default:
//			fmt.Println("请输入正确的选项...")
//		}
//
//		if !loop {
//			break
//		}
//	}
//	fmt.Println("你退出家庭记账软件的使用...")
//}
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	//声明一个变量，保存接收用户输入的选项
//	key := ""
//	//声明一个变量，控制是否退出for
//	loop := true
//
//	//定义账户的余额 []
//	balance := 10000.0
//	//每次收支的金额
//	money := 0.0
//	//每次收支的说明
//	note := ""
//	//定义个变量，记录是否有收支的行为
//	flag := false
//	//收支的详情使用字符串来记录
//	//当有收支时，只需要对details 进行拼接处理即可
//	details := "收支\t账户金额\t收支金额\t说    明"
//
//	//显示这个主菜单
//	for {
//		fmt.Println("\n-----------------家庭收支记账软件-----------------")
//		fmt.Println("                  1 收支明细")
//		fmt.Println("                  2 登记收入")
//		fmt.Println("                  3 登记支出")
//		fmt.Println("                  4 退出软件")
//		fmt.Print("请选择(1-4)：")
//		fmt.Scanln(&key)
//
//		switch key {
//		case "1":
//			fmt.Println("-----------------当前收支明细记录-----------------")
//			if flag {
//				fmt.Println(details)
//			} else {
//				fmt.Println("当前没有收支明细... 来一笔吧!")
//			}
//		case "2":
//			fmt.Println("本次收入金额:")
//			fmt.Scanln(&money)
//			balance += money // 修改账户余额
//			fmt.Println("本次收入说明:")
//			fmt.Scanln(&note)
//			//将这个收入情况，拼接到details变量
//			//收入    11000           1000            有人发红包
//			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
//		case "3":
//			fmt.Println("本次支出金额:")
//			fmt.Scanln(&money)
//
//			balance -= money
//			fmt.Println("本次支出说明:")
//			fmt.Scanln(&note)
//			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
//		case "4":
//			fmt.Println("你确定要退出吗? y/n")
//			choice := ""
//			for {
//
//				fmt.Scanln(&choice)
//				if choice == "y" || choice == "n" {
//					break
//				}
//				fmt.Println("你的输入有误，请重新输入 y/n")
//			}
//
//			if choice == "y" {
//				loop = false
//			}
//		default:
//			fmt.Println("请输入正确的选项...")
//		}
//
//		if !loop {
//			break
//		}
//	}
//	fmt.Println("你退出家庭记账软件的使用...")
//}
