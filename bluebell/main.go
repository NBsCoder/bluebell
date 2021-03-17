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
//package utils
//import (
//"fmt"
//)
//
//type FamilyAccount struct {
//	//声明必须的字段.
//
//	//声明一个字段，保存接收用户输入的选项
//	key  string
//	//声明一个字段，控制是否退出for
//	loop bool
//	//定义账户的余额 []
//	balance float64
//	//每次收支的金额
//	money float64
//	//每次收支的说明
//	note string
//	//定义个字段，记录是否有收支的行为
//	flag bool
//	//收支的详情使用字符串来记录
//	//当有收支时，只需要对details 进行拼接处理即可
//	details string
//}
//
////编写要给工厂模式的构造方法，返回一个*FamilyAccount实例
//func NewFamilyAccount() *FamilyAccount {
//
//	return &FamilyAccount{
//		key : "",
//		loop : true,
//		balance : 10000.0,
//		money : 0.0,
//		note : "",
//		flag : false,
//		details : "收支\t账户金额\t收支金额\t说    明",
//	}
//
//}
//
////将显示明细写成一个方法
//func (this *FamilyAccount) showDetails() {
//	fmt.Println("-----------------当前收支明细记录-----------------")
//	if this.flag {
//		fmt.Println(this.details)
//	} else {
//		fmt.Println("当前没有收支明细... 来一笔吧!")
//	}
//}
//
////将登记收入写成一个方法，和*FamilyAccount绑定
//func (this *FamilyAccount) income() {
//
//	fmt.Println("本次收入金额:")
//	fmt.Scanln(&this.money)
//	this.balance += this.money // 修改账户余额
//	fmt.Println("本次收入说明:")
//	fmt.Scanln(&this.note)
//	//将这个收入情况，拼接到details变量
//	//收入    11000           1000            有人发红包
//	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
//	this.flag = true
//}
//
////将登记支出写成一个方法，和*FamilyAccount绑定
//func (this *FamilyAccount) pay() {
//	fmt.Println("本次支出金额:")
//	fmt.Scanln(&this.money)
//	//这里需要做一个必要的判断
//	if this.money > this.balance {
//		fmt.Println("余额的金额不足")
//		//break
//	}
//	this.balance -= this.money
//	fmt.Println("本次支出说明:")
//	fmt.Scanln(&this.note)
//	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
//	this.flag = true
//}
//
////将退出系统写成一个方法,和*FamilyAccount绑定
//func (this *FamilyAccount) exit() {
//
//	fmt.Println("你确定要退出吗? y/n")
//	choice := ""
//	for {
//
//		fmt.Scanln(&choice)
//		if choice == "y" || choice == "n" {
//			break
//		}
//		fmt.Println("你的输入有误，请重新输入 y/n")
//	}
//
//	if choice == "y" {
//		this.loop = false
//	}
//}
//
//
////给该结构体绑定相应的方法
////显示主菜单
//func (this *FamilyAccount) MainMenu() {
//
//	for {
//		fmt.Println("\n-----------------家庭收支记账软件-----------------")
//		fmt.Println("                  1 收支明细")
//		fmt.Println("                  2 登记收入")
//		fmt.Println("                  3 登记支出")
//		fmt.Println("                  4 退出软件")
//		fmt.Print("请选择(1-4)：")
//		fmt.Scanln(&this.key)
//		switch this.key {
//		case "1":
//			this.showDetails()
//		case "2":
//			this.income()
//		case "3":
//			this.pay()
//		case "4":
//			this.exit()
//		default :
//			fmt.Println("请输入正确的选项..")
//		}
//
//		if !this.loop {
//			break
//		}
//
//	}
//}
//package model
//import (
//"fmt"
//)
////声明一个Customer结构体，表示一个客户信息
//
//type Customer struct {
//	Id int
//	Name string
//	Gender string
//	Age int
//	Phone string
//	Email string
//}
//
////使用工厂模式，返回一个Customer的实例
//
//func NewCustomer(id int, name string, gender string,
//	age int, phone string, email string ) Customer {
//	return Customer{
//		Id : id,
//		Name : name,
//		Gender : gender,
//		Age : age,
//		Phone : phone,
//		Email : email,
//	}
//}
//
////第二种创建Customer实例方法，不带id
//func NewCustomer2(name string, gender string,
//	age int, phone string, email string ) Customer {
//	return Customer{
//		Name : name,
//		Gender : gender,
//		Age : age,
//		Phone : phone,
//		Email : email,
//	}
//}
//
////返回用户的信息,格式化的字符串
//func (this Customer) GetInfo()  string {
//	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
//		this.Name, this.Gender,this.Age, this.Phone, this.Email)
//	return info
//
//}
//package service
//import (
//"go_code/customerManage/model"
//
//)
//
////该CustomerService， 完成对Customer的操作,包括
////增删改查
//type CustomerService struct {
//	customers []model.Customer
//	//声明一个字段，表示当前切片含有多少个客户
//	//该字段后面，还可以作为新客户的id+1
//	customerNum int
//}
//
////编写一个方法，可以返回 *CustomerService
//func NewCustomerService() *CustomerService {
//	//为了能够看到有客户在切片中，我们初始化一个客户
//	customerService := &CustomerService{}
//	customerService.customerNum = 1
//	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
//	customerService.customers = append(customerService.customers, customer)
//	return customerService
//}
//
////返回客户切片
//func (this *CustomerService) List() []model.Customer {
//	return this.customers
//}
//
////添加客户到customers切片
////!!!
//func (this *CustomerService) Add(customer model.Customer) bool {
//
//	//我们确定一个分配id的规则,就是添加的顺序
//	this.customerNum++
//	customer.Id = this.customerNum
//	this.customers = append(this.customers, customer)
//	return true
//}
//
////根据id删除客户(从切片中删除)
//func (this *CustomerService) Delete(id int) bool {
//	index := this.FindById(id)
//	//如果index == -1, 说明没有这个客户
//	if index == -1 {
//		return false
//	}
//	//如何从切片中删除一个元素
//	this.customers = append(this.customers[:index], this.customers[index+1:]...)
//	return true
//}
//
////根据id查找客户在切片中对应下标,如果没有该客户，返回-1
//func (this *CustomerService) FindById(id int)  int {
//	index := -1
//	//遍历this.customers 切片
//	for i := 0; i < len(this.customers); i++ {
//		if this.customers[i].Id == id {
//			//找到
//			index = i
//		}
//	}
//	return index
//}
//package main
//
//import (
//"fmt"
//"go_code/customerManage/service"
//"go_code/customerManage/model"
//)
//
//type customerView struct {
//
//	//定义必要字段
//	key string //接收用户输入...
//	loop bool  //表示是否循环的显示主菜单
//	//增加一个字段customerService
//	customerService *service.CustomerService
//
//}
//
////显示所有的客户信息
//func (this *customerView) list() {
//
//	//首先，获取到当前所有的客户信息(在切片中)
//	customers := this.customerService.List()
//	//显示
//	fmt.Println("---------------------------客户列表---------------------------")
//	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
//	for i := 0; i < len(customers); i++ {
//		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
//		fmt.Println(customers[i].GetInfo())
//	}
//	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
//}
//
////得到用户的输入，信息构建新的客户，并完成添加
//func (this *customerView) add() {
//	fmt.Println("---------------------添加客户---------------------")
//	fmt.Println("姓名:")
//	name := ""
//	fmt.Scanln(&name)
//	fmt.Println("性别:")
//	gender := ""
//	fmt.Scanln(&gender)
//	fmt.Println("年龄:")
//	age := 0
//	fmt.Scanln(&age)
//	fmt.Println("电话:")
//	phone := ""
//	fmt.Scanln(&phone)
//	fmt.Println("电邮:")
//	email := ""
//	fmt.Scanln(&email)
//	//构建一个新的Customer实例
//	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
//	customer := model.NewCustomer2(name, gender, age, phone, email)
//	//调用
//	if this.customerService.Add(customer) {
//		fmt.Println("---------------------添加完成---------------------")
//	} else {
//		fmt.Println("---------------------添加失败---------------------")
//	}
//}
//
////得到用户的输入id，删除该id对应的客户
//func (this *customerView) delete() {
//	fmt.Println("---------------------删除客户---------------------")
//	fmt.Println("请选择待删除客户编号(-1退出)：")
//	id := -1
//	fmt.Scanln(&id)
//	if id == -1 {
//		return //放弃删除操作
//	}
//	fmt.Println("确认是否删除(Y/N)：")
//	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
//	choice := ""
//	fmt.Scanln(&choice)
//	if choice == "y" || choice == "Y" {
//		//调用customerService 的 Delete方法
//		if this.customerService.Delete(id) {
//			fmt.Println("---------------------删除完成---------------------")
//		} else {
//			fmt.Println("---------------------删除失败，输入的id号不存在----")
//		}
//	}
//}
//
////退出软件
//func (this *customerView) exit() {
//
//	fmt.Println("确认是否退出(Y/N)：")
//	for {
//		fmt.Scanln(&this.key)
//		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
//			break
//		}
//
//		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
//	}
//
//	if this.key == "Y" || this.key == "y" {
//		this.loop = false
//	}
//
//}
//
////显示主菜单
//func (this *customerView) mainMenu() {
//
//	for {
//
//		fmt.Println("-----------------客户信息管理软件-----------------")
//		fmt.Println("                 1 添 加 客 户")
//		fmt.Println("                 2 修 改 客 户")
//		fmt.Println("                 3 删 除 客 户")
//		fmt.Println("                 4 客 户 列 表")
//		fmt.Println("                 5 退       出")
//		fmt.Print("请选择(1-5)：")
//
//		fmt.Scanln(&this.key)
//		switch this.key {
//		case "1" :
//			this.add()
//		case "2" :
//			fmt.Println("修 改 客 户")
//		case "3" :
//			this.delete()
//		case "4" :
//			this.list()
//		case "5" :
//			this.exit()
//		default :
//			fmt.Println("你的输入有误，请重新输入...")
//		}
//
//		if !this.loop {
//			break
//		}
//
//	}
//	fmt.Println("你退出了客户关系管理系统...")
//}
//
//
//
//func main() {
//	//在main函数中，创建一个customerView,并运行显示主菜单..
//	customerView := customerView{
//		key : "",
//		loop : true,
//	}
//	//这里完成对customerView结构体的customerService字段的初始化
//	customerView.customerService = service.NewCustomerService()
//	//显示主菜单..
//	customerView.mainMenu()
//
//}
//package model
//import (
//"fmt"
//)
////声明一个Customer结构体，表示一个客户信息
//
//type Customer struct {
//	Id int
//	Name string
//	Gender string
//	Age int
//	Phone string
//	Email string
//}
//
////使用工厂模式，返回一个Customer的实例
//
//func NewCustomer(id int, name string, gender string,
//	age int, phone string, email string ) Customer {
//	return Customer{
//		Id : id,
//		Name : name,
//		Gender : gender,
//		Age : age,
//		Phone : phone,
//		Email : email,
//	}
//}
//
////第二种创建Customer实例方法，不带id
//func NewCustomer2(name string, gender string,
//	age int, phone string, email string ) Customer {
//	return Customer{
//		Name : name,
//		Gender : gender,
//		Age : age,
//		Phone : phone,
//		Email : email,
//	}
//}
//
////返回用户的信息,格式化的字符串
//func (this Customer) GetInfo()  string {
//	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
//		this.Name, this.Gender,this.Age, this.Phone, this.Email)
//	return info
//
//}
//package service
//import (
//"go_code/customerManage/model"
//
//)
//
////该CustomerService， 完成对Customer的操作,包括
////增删改查
//type CustomerService struct {
//	customers []model.Customer
//	//声明一个字段，表示当前切片含有多少个客户
//	//该字段后面，还可以作为新客户的id+1
//	customerNum int
//}
//
////编写一个方法，可以返回 *CustomerService
//func NewCustomerService() *CustomerService {
//	//为了能够看到有客户在切片中，我们初始化一个客户
//	customerService := &CustomerService{}
//	customerService.customerNum = 1
//	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
//	customerService.customers = append(customerService.customers, customer)
//	return customerService
//}
//
////返回客户切片
//func (this *CustomerService) List() []model.Customer {
//	return this.customers
//}
//
////添加客户到customers切片
////!!!
//func (this *CustomerService) Add(customer model.Customer) bool {
//
//	//我们确定一个分配id的规则,就是添加的顺序
//	this.customerNum++
//	customer.Id = this.customerNum
//	this.customers = append(this.customers, customer)
//	return true
//}
//
////根据id删除客户(从切片中删除)
//func (this *CustomerService) Delete(id int) bool {
//	index := this.FindById(id)
//	//如果index == -1, 说明没有这个客户
//	if index == -1 {
//		return false
//	}
//	//如何从切片中删除一个元素
//	this.customers = append(this.customers[:index], this.customers[index+1:]...)
//	return true
//}
//
////根据id查找客户在切片中对应下标,如果没有该客户，返回-1
//func (this *CustomerService) FindById(id int)  int {
//	index := -1
//	//遍历this.customers 切片
//	for i := 0; i < len(this.customers); i++ {
//		if this.customers[i].Id == id {
//			//找到
//			index = i
//		}
//	}
//	return index
//}
//package main
//
//import (
//"fmt"
//"go_code/customerManage/service"
//"go_code/customerManage/model"
//)
//
//type customerView struct {
//
//	//定义必要字段
//	key string //接收用户输入...
//	loop bool  //表示是否循环的显示主菜单
//	//增加一个字段customerService
//	customerService *service.CustomerService
//
//}
//
////显示所有的客户信息
//func (this *customerView) list() {
//
//	//首先，获取到当前所有的客户信息(在切片中)
//	customers := this.customerService.List()
//	//显示
//	fmt.Println("---------------------------客户列表---------------------------")
//	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
//	for i := 0; i < len(customers); i++ {
//		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
//		fmt.Println(customers[i].GetInfo())
//	}
//	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
//}
//
////得到用户的输入，信息构建新的客户，并完成添加
//func (this *customerView) add() {
//	fmt.Println("---------------------添加客户---------------------")
//	fmt.Println("姓名:")
//	name := ""
//	fmt.Scanln(&name)
//	fmt.Println("性别:")
//	gender := ""
//	fmt.Scanln(&gender)
//	fmt.Println("年龄:")
//	age := 0
//	fmt.Scanln(&age)
//	fmt.Println("电话:")
//	phone := ""
//	fmt.Scanln(&phone)
//	fmt.Println("电邮:")
//	email := ""
//	fmt.Scanln(&email)
//	//构建一个新的Customer实例
//	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
//	customer := model.NewCustomer2(name, gender, age, phone, email)
//	//调用
//	if this.customerService.Add(customer) {
//		fmt.Println("---------------------添加完成---------------------")
//	} else {
//		fmt.Println("---------------------添加失败---------------------")
//	}
//}
//
////得到用户的输入id，删除该id对应的客户
//func (this *customerView) delete() {
//	fmt.Println("---------------------删除客户---------------------")
//	fmt.Println("请选择待删除客户编号(-1退出)：")
//	id := -1
//	fmt.Scanln(&id)
//	if id == -1 {
//		return //放弃删除操作
//	}
//	fmt.Println("确认是否删除(Y/N)：")
//	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
//	choice := ""
//	fmt.Scanln(&choice)
//	if choice == "y" || choice == "Y" {
//		//调用customerService 的 Delete方法
//		if this.customerService.Delete(id) {
//			fmt.Println("---------------------删除完成---------------------")
//		} else {
//			fmt.Println("---------------------删除失败，输入的id号不存在----")
//		}
//	}
//}
//
////退出软件
//func (this *customerView) exit() {
//
//	fmt.Println("确认是否退出(Y/N)：")
//	for {
//		fmt.Scanln(&this.key)
//		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
//			break
//		}
//
//		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
//	}
//
//	if this.key == "Y" || this.key == "y" {
//		this.loop = false
//	}
//
//}
//
////显示主菜单
//func (this *customerView) mainMenu() {
//
//	for {
//
//		fmt.Println("-----------------客户信息管理软件-----------------")
//		fmt.Println("                 1 添 加 客 户")
//		fmt.Println("                 2 修 改 客 户")
//		fmt.Println("                 3 删 除 客 户")
//		fmt.Println("                 4 客 户 列 表")
//		fmt.Println("                 5 退       出")
//		fmt.Print("请选择(1-5)：")
//
//		fmt.Scanln(&this.key)
//		switch this.key {
//		case "1" :
//			this.add()
//		case "2" :
//			fmt.Println("修 改 客 户")
//		case "3" :
//			this.delete()
//		case "4" :
//			this.list()
//		case "5" :
//			this.exit()
//		default :
//			fmt.Println("你的输入有误，请重新输入...")
//		}
//
//		if !this.loop {
//			break
//		}
//
//	}
//	fmt.Println("你退出了客户关系管理系统...")
//}
//
//
//
//func main() {
//	//在main函数中，创建一个customerView,并运行显示主菜单..
//	customerView := customerView{
//		key : "",
//		loop : true,
//	}
//	//这里完成对customerView结构体的customerService字段的初始化
//	customerView.customerService = service.NewCustomerService()
//	//显示主菜单..
//	customerView.mainMenu()
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
//package main
//import (
//"fmt"
//"os"
//"io"
//"bufio"
//)
//
////定义一个结构体，用于保存统计结果
//type CharCount struct {
//	ChCount int // 记录英文个数
//	NumCount int // 记录数字的个数
//	SpaceCount int // 记录空格的个数
//	OtherCount int // 记录其它字符的个数
//}
//
//func main() {
//
//	//思路: 打开一个文件, 创一个Reader
//	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
//	//然后将结果保存到一个结构体
//	fileName := "e:/abc.txt"
//	file, err := os.Open(fileName)
//	if err != nil {
//		fmt.Printf("open file err=%v\n", err)
//		return
//	}
//	defer file.Close()
//	//定义个CharCount 实例
//	var count CharCount
//	//创建一个Reader
//	reader := bufio.NewReader(file)
//
//	//开始循环的读取fileName的内容
//	for {
//		str, err := reader.ReadString('\n')
//		if err == io.EOF { //读到文件末尾就退出
//			break
//		}
//		//遍历 str ，进行统计
//		for _, v := range str {
//
//			switch {
//			case v >= 'a' && v <= 'z':
//				fallthrough //穿透
//			case v >= 'A' && v <= 'Z':
//				count.ChCount++
//			case v == ' ' || v == '\t':
//				count.SpaceCount++
//			case v >= '0' && v <= '9':
//				count.NumCount++
//			default :
//				count.OtherCount++
//			}
//		}
//	}
//
//	//输出统计的结果看看是否正确
//	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
//		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
//
//}
//package main
//import (
//"fmt"
//"flag"
//)
//
//func main() {
//
//	//定义几个变量，用于接收命令行的参数值
//	var user string
//	var pwd string
//	var host string
//	var port int
//
//	//&user 就是接收用户命令行中输入的 -u 后面的参数值
//	//"u" ,就是 -u 指定参数
//	//"" , 默认值
//	//"用户名,默认为空" 说明
//	flag.StringVar(&user, "u", "", "用户名,默认为空")
//	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
//	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
//	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
//	//这里有一个非常重要的操作,转换， 必须调用该方法
//	flag.Parse()
//
//	//输出结果
//	fmt.Printf("user=%v pwd=%v host=%v port=%v",
//		user, pwd, host, port)
//
//}
//package main
//import (
//"fmt"
//"encoding/json"
//)
//
////定义一个结构体
//type Monster struct {
//	Name string `json:"monster_name"` //反射机制
//	Age int `json:"monster_age"`
//	Birthday string //....
//	Sal float64
//	Skill string
//}
//
//func testStruct() {
//	//演示
//	monster := Monster{
//		Name :"牛魔王",
//		Age : 500 ,
//		Birthday : "2011-11-11",
//		Sal : 8000.0,
//		Skill : "牛魔拳",
//	}
//
//	//将monster 序列化
//	data, err := json.Marshal(&monster) //..
//	if err != nil {
//		fmt.Printf("序列号错误 err=%v\n", err)
//	}
//	//输出序列化后的结果
//	fmt.Printf("monster序列化后=%v\n", string(data))
//
//}
//
////将map进行序列化
//func testMap() {
//	//定义一个map
//	var a map[string]interface{}
//	//使用map,需要make
//	a = make(map[string]interface{})
//	a["name"] = "红孩儿"
//	a["age"] = 30
//	a["address"] = "洪崖洞"
//
//	//将a这个map进行序列化
//	//将monster 序列化
//	data, err := json.Marshal(a)
//	if err != nil {
//		fmt.Printf("序列化错误 err=%v\n", err)
//	}
//	//输出序列化后的结果
//	fmt.Printf("a map 序列化后=%v\n", string(data))
//
//}
//
////演示对切片进行序列化, 我们这个切片 []map[string]interface{}
//func testSlice() {
//	var slice []map[string]interface{}
//	var m1 map[string]interface{}
//	//使用map前，需要先make
//	m1 = make(map[string]interface{})
//	m1["name"] = "jack"
//	m1["age"] = "7"
//	m1["address"] = "北京"
//	slice = append(slice, m1)
//
//	var m2 map[string]interface{}
//	//使用map前，需要先make
//	m2 = make(map[string]interface{})
//	m2["name"] = "tom"
//	m2["age"] = "20"
//	m2["address"] = [2]string{"墨西哥","夏威夷"}
//	slice = append(slice, m2)
//
//	//将切片进行序列化操作
//	data, err := json.Marshal(slice)
//	if err != nil {
//		fmt.Printf("序列化错误 err=%v\n", err)
//	}
//	//输出序列化后的结果
//	fmt.Printf("slice 序列化后=%v\n", string(data))
//
//}
//
////对基本数据类型序列化，对基本数据类型进行序列化意义不大
//func testFloat64() {
//	var num1 float64 = 2345.67
//
//	//对num1进行序列化
//	data, err := json.Marshal(num1)
//	if err != nil {
//		fmt.Printf("序列化错误 err=%v\n", err)
//	}
//	//输出序列化后的结果
//	fmt.Printf("num1 序列化后=%v\n", string(data))
//}
//
//func main() {
//	//演示将结构体, map , 切片进行序列号
//	testStruct()
//	testMap()
//	testSlice()//演示对切片的序列化
//	testFloat64()//演示对基本数据类型的序列化
//}
//package main
//import "fmt"
///*
//	arr=[1,1,1,2,2,3,4,3]
//	编写一个函数，输入数组，输出数组中重复最多的元素，及对应重复次数
//	返回出现最多的元素及其出现次数
//*/
//var b map[int]int
///*
//	1、将不同的元素存为map的key
//	2、拿key去数组中遍历，出现一次，计数加1
//	3、将次数存为map的值
//	4、比较map的值，选出最大值并将键一起打印出来
//*/
//func f(a []int) map[int]int{
//	var s []int
//	for i := 0; i < len(a)-1; i++ { //1、遍历拿到map的key,存到一个切片中
//		if a[i] != a[i+1] {
//			s = append(s, a[i])
//		}
//	}
//	for _, v := range s {			//2.遍历拿到map的value
//		var count int = 0
//		for _, y := range a {
//			if v == y {
//				count++
//			}
//			b[v] = count
//		}
//	}
//	for m, _ := range b {			//3.冒泡排序
//		fmt.Printf("数字%d出现了%d次\n", m, b[m])
//	}
//	return b
//}
//func main() {
//	var arr = []int{1, 1, 1, 2,2, 2, 3, 4, 3}
//	b = make(map[int]int)
//	v:=f(arr)
//	var max int
//	var maxKey []int
//	for k,w:=range v{
//		if w>max||max==0{
//			max=w
//		}
//		if max==v[k]{
//			maxKey=append(maxKey,k)
//		}
//	}
//	fmt.Printf("出现最多的数字是%d，最多的数字%d次\n",maxKey,max)
//}
//package main
//import (
//"fmt"
//"time"
//)
//
//// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
//// 最后显示出来。要求使用goroutine完成
//
//// 思路
//// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
//// 2. 我们启动的协程多个，统计的将结果放入到 map中
//// 3. map 应该做出一个全局的.
//
//var (
//	myMap = make(map[int]int, 10)
//)
//
//// test 函数就是计算 n!, 让将这个结果放入到 myMap
//func test(n int) {
//
//	res := 1
//	for i := 1; i <= n; i++ {
//		res *= i
//	}
//
//	//这里我们将 res 放入到myMap
//	myMap[n] = res //concurrent map writes?
//}
//
//func main() {
//
//	// 我们这里开启多个协程完成这个任务[200个]
//	for i := 1; i <= 200; i++ {
//		go test(i)
//	}
//
//
//	//休眠10秒钟【第二个问题 】
//	time.Sleep(time.Second * 10)
//
//	//这里我们输出结果,变量这个结果
//	for i, v := range myMap {
//		fmt.Printf("map[%d]=%d\n", i, v)
//	}
//}
//package main
//import (
//"fmt"
//_ "time"
//"sync"
//)
//
//// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
//// 最后显示出来。要求使用goroutine完成
//
//// 思路
//// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
//// 2. 我们启动的协程多个，统计的将结果放入到 map中
//// 3. map 应该做出一个全局的.
//
//var (
//	myMap = make(map[int]int, 10)
//	//声明一个全局的互斥锁
//	//lock 是一个全局的互斥锁，
//	//sync 是包: synchornized 同步
//	//Mutex : 是互斥
//	lock sync.Mutex
//)
//
//// test 函数就是计算 n!, 让将这个结果放入到 myMap
//func test(n int) {
//
//	res := 1
//	for i := 1; i <= n; i++ {
//		res *= i
//	}
//
//	//这里我们将 res 放入到myMap
//	//加锁
//	lock.Lock()
//	myMap[n] = res //concurrent map writes?
//	//解锁
//	lock.Unlock()
//}
//
//func main() {
//
//	// 我们这里开启多个协程完成这个任务[200个]
//	for i := 1; i <= 20; i++ {
//		go test(i)
//	}
//
//	//休眠10秒钟【第二个问题 】
//	//time.Sleep(time.Second * 5)
//
//	//这里我们输出结果,变量这个结果
//	lock.Lock()
//	for i, v := range myMap {
//		fmt.Printf("map[%d]=%d\n", i, v)
//	}
//	lock.Unlock()
//
//}
//package main
//import (
//"fmt"
//)
//
////如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间
////如果需要使用这样的字段，需要先make，才能使用.
//
//type Person struct{
//	Name string
//	Age int
//	Scores [5]float64
//	ptr *int //指针
//	slice []int //切片
//	map1 map[string]string //map
//}
//
//func main() {
//
//	//定义结构体变量
//	var p1 Person
//	fmt.Println(p1)
//
//	if p1.ptr == nil {
//		fmt.Println("ok1")
//	}
//
//	if p1.slice == nil {
//		fmt.Println("ok2")
//	}
//
//	if p1.map1 == nil {
//		fmt.Println("ok3")
//	}
//
//	//使用slice, 再次说明，一定要make
//	p1.slice = make([]int, 10)
//	p1.slice[0] = 100 //ok
//
//	//使用map, 一定要先make
//	p1.map1 = make(map[string]string)
//	p1.map1["key1"] = "tom~"
//	fmt.Println(p1)
//
//}
//package main
//
//import (
//"fmt"
//"go_code/customerManage/service"
//"go_code/customerManage/model"
//)
//
//type customerView struct {
//
//	//定义必要字段
//	key string //接收用户输入...
//	loop bool  //表示是否循环的显示主菜单
//	//增加一个字段customerService
//	customerService *service.CustomerService
//
//}
//
////显示所有的客户信息
//func (this *customerView) list() {
//
//	//首先，获取到当前所有的客户信息(在切片中)
//	customers := this.customerService.List()
//	//显示
//	fmt.Println("---------------------------客户列表---------------------------")
//	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
//	for i := 0; i < len(customers); i++ {
//		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
//		fmt.Println(customers[i].GetInfo())
//	}
//	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
//}
//
////得到用户的输入，信息构建新的客户，并完成添加
//func (this *customerView) add() {
//	fmt.Println("---------------------添加客户---------------------")
//	fmt.Println("姓名:")
//	name := ""
//	fmt.Scanln(&name)
//	fmt.Println("性别:")
//	gender := ""
//	fmt.Scanln(&gender)
//	fmt.Println("年龄:")
//	age := 0
//	fmt.Scanln(&age)
//	fmt.Println("电话:")
//	phone := ""
//	fmt.Scanln(&phone)
//	fmt.Println("电邮:")
//	email := ""
//	fmt.Scanln(&email)
//	//构建一个新的Customer实例
//	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
//	customer := model.NewCustomer2(name, gender, age, phone, email)
//	//调用
//	if this.customerService.Add(customer) {
//		fmt.Println("---------------------添加完成---------------------")
//	} else {
//		fmt.Println("---------------------添加失败---------------------")
//	}
//}
//
////得到用户的输入id，删除该id对应的客户
//func (this *customerView) delete() {
//	fmt.Println("---------------------删除客户---------------------")
//	fmt.Println("请选择待删除客户编号(-1退出)：")
//	id := -1
//	fmt.Scanln(&id)
//	if id == -1 {
//		return //放弃删除操作
//	}
//	fmt.Println("确认是否删除(Y/N)：")
//	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
//	choice := ""
//	fmt.Scanln(&choice)
//	if choice == "y" || choice == "Y" {
//		//调用customerService 的 Delete方法
//		if this.customerService.Delete(id) {
//			fmt.Println("---------------------删除完成---------------------")
//		} else {
//			fmt.Println("---------------------删除失败，输入的id号不存在----")
//		}
//	}
//}
//
////退出软件
//func (this *customerView) exit() {
//
//	fmt.Println("确认是否退出(Y/N)：")
//	for {
//		fmt.Scanln(&this.key)
//		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
//			break
//		}
//
//		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
//	}
//
//	if this.key == "Y" || this.key == "y" {
//		this.loop = false
//	}
//
//}
//
////显示主菜单
//func (this *customerView) mainMenu() {
//
//	for {
//
//		fmt.Println("-----------------客户信息管理软件-----------------")
//		fmt.Println("                 1 添 加 客 户")
//		fmt.Println("                 2 修 改 客 户")
//		fmt.Println("                 3 删 除 客 户")
//		fmt.Println("                 4 客 户 列 表")
//		fmt.Println("                 5 退       出")
//		fmt.Print("请选择(1-5)：")
//
//		fmt.Scanln(&this.key)
//		switch this.key {
//		case "1" :
//			this.add()
//		case "2" :
//			fmt.Println("修 改 客 户")
//		case "3" :
//			this.delete()
//		case "4" :
//			this.list()
//		case "5" :
//			this.exit()
//		default :
//			fmt.Println("你的输入有误，请重新输入...")
//		}
//
//		if !this.loop {
//			break
//		}
//
//	}
//	fmt.Println("你退出了客户关系管理系统...")
//}
//
//
//
//func main() {
//	//在main函数中，创建一个customerView,并运行显示主菜单..
//	customerView := customerView{
//		key : "",
//		loop : true,
//	}
//	//这里完成对customerView结构体的customerService字段的初始化
//	customerView.customerService = service.NewCustomerService()
//	//显示主菜单..
//	customerView.mainMenu()
//
//}
//package main
//import (
//"fmt"
//"os"
//"bufio"
//"io"
//)
//func main() {
//	//打开文件
//	//概念说明: file 的叫法
//	//1. file 叫 file对象
//	//2. file 叫 file指针
//	//3. file 叫 file 文件句柄
//	file , err := os.Open("d:/test.txt")
//	if err != nil {
//		fmt.Println("open file err=", err)
//	}
//
//	//当函数退出时，要及时的关闭file
//	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏.
//
//	// 创建一个 *Reader  ，是带缓冲的
//	/*
//		const (
//		defaultBufSize = 4096 //默认的缓冲区为4096
//		)
//	*/
//	reader := bufio.NewReader(file)
//	//循环的读取文件的内容
//	for {
//		str, err := reader.ReadString('\n') // 读到一个换行就结束
//		if err == io.EOF { // io.EOF表示文件的末尾
//			break
//		}
//		//输出内容
//		fmt.Printf(str)
//	}
//
//	fmt.Println("文件读取结束...")
//}
////package main
////
////import (
////	"fmt"
////	"time"
////)
//
/////*
////	arr=[1,1,1,2,2,3,4,3]
////	编写一个函数，输入数组，输出数组中重复最多的元素，及对应重复次数
////	返回出现最多的元素及其出现次数
////*/
////var b map[int]int
/////*
////	1、将不同的元素存为map的key
////	2、拿key去数组中遍历，出现一次，计数加1
////	3、将次数存为map的值
////	4、比较map的值，选出最大值并将键一起打印出来
////*/
////func f(a []int) map[int]int{
////	var s []int
////	for i := 0; i < len(a)-1; i++ { //1、遍历拿到map的key,存到一个切片中
////		if a[i] != a[i+1] {
////			s = append(s, a[i])
////		}
////	}
////	for _, v := range s {			//2.遍历拿到map的value
////		var count int = 0
////		for _, y := range a {
////			if v == y {
////				count++
////			}
////			b[v] = count
////		}
////	}
////	for m, _ := range b {			//3.冒泡排序
////		fmt.Printf("数字%d出现了%d次\n", m, b[m])
////	}
////	return b
////}
////func main() {
////	var arr = []int{1,1,1, 1, 1, 2,2, 2,2,2, 3, 4, 3,3}
////	b = make(map[int]int)
////	v:=f(arr)
////	var max int
////	var maxKey []int
////	for k,w:=range v{
////		if w>max||max==0{
////			max=w
////		}
////		if max==v[k]{
////			maxKey=append(maxKey,k)
////		}
////	}
////	fmt.Printf("出现最多的数字是%d，最多的数字%d次\n",maxKey,max)
////}
//
//// 1. s = []string{"a","cc","sa","cc"} 需求: 切片字符串去重 期望返回值: [ "a" "cc" "sa"]
//// 2. s = []string{"a","cc","sa","cc"} 传入: "cc" 需求: 根据传入值,删除切片元素  期望得到: ["a" "sa"]
////func main(){
////	var s =[]string{"a","cc","sa","cc"}
////	m:= make(map[string]string)
////	var b []string
////	for _,value:=range s{
////		m[value]="a"
////	}
////	for k,_:=range m{
////		b=append(b,k)
////	}
////	fmt.Println(b)
////}
//
////请判断"abcdcba"字符串是否是回文
////
////package main
////
////import (
////	"fmt"
////)
////
//////write Data
////func writeData(intChan chan int) {
////	for i := 1; i <= 50; i++ {
////		//放入数据
////		intChan <- i //
////		fmt.Println("writeData ", i)
////		//time.Sleep(time.Second)
////	}
////	close(intChan) //关闭
////}
////
//////read data
////func readData(intChan chan int,exitChan chan bool) {
////
////	for {
////		v, ok := <-intChan
////		if !ok {
////			break
////		}
////		//time.Sleep(time.Second)
////		fmt.Printf("readData 读到数据=%v\n", v)
////	}
////	//readData 读取完数据后，即任务完成
////	exitChan <- true
////	close(exitChan)
////
////}
////
////func main() {
////
////	//创建两个管道
////	intChan := make(chan int, 10)
////	exitChan := make(chan bool, 1)
////	go writeData(intChan)
////	go readData(intChan,exitChan)
////	for {
////		value, ok := <-exitChan
////		if !ok {
////			break
////		}
////		fmt.Println(value)
////	}
////
////}
//package main
//
//import "fmt"
//
////往numChan中写入2000个数字
//func writeDate(numChan chan int){
//	for i:=1;i<=2000;i++{
//		numChan<-i
//	}
////写完了，关闭通道
//	close(numChan)
//}
////从numChan中取出数字并计算各个数字的累加
//func addSum(numChan chan int,resChan chan int64,exitChan chan bool)  {
//	var sum int
//	for v:=range numChan{
//		sum +=v
//		resChan<-int64(sum)
//	}
//	close(resChan)
//	exitChan<-true
//}
//func main()  {
//	a:=2.0
//	fmt.Println(a)
//	 numChan:=make(chan int,1000)
//	 resChan:=make(chan int64,2000)
//	 exitChan:=make(chan bool,8)
//	 //1、开启一个协程往numChan中加数字
//	go writeDate(numChan)
//	 //2、开启8个协程，读取numChan中的值并累加，存到resChan中
//	 for i:=0;i<8;i++{
//	 	go addSum(numChan,resChan,exitChan)
//	 }
//	 //3、遍历exitChan，等待协程运算完毕
//	go func() {
//		//能读到8个值
//		for j:=0;j<8;j++{
//			<-exitChan
//		}
//		//读完8个值就关闭这个通道
//		close(exitChan)
//	}()
//	 //4、遍历resChan通道，并将结果处理一下打印出来
//
//}
//package main
//
//import "fmt"

//func main()  {
//	ch:='b'					//int32
//	fmt.Println(2.00000/ch)		//int32 / float64	=	49
//	//ch:='b'
//	//a:=2.0
//	//fmt.Println(ch/a)		这里编译时会报错
//}
//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	return 5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	return 5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}
//type A interface {
//	ShowA() int
//}
//
//type B interface {
//	ShowB() int
//}
//
//type Work struct {
//	i int
//}
//
//func (w Work) ShowA() int {
//	return w.i + 10
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main() {
//	var a A = Work{3}
//	s := a
//	fmt.Println(s.ShowA())
//	fmt.Println(s.ShowB())
//}
//type A interface {
//	ShowA() int
//}
//
//type B interface {
//	ShowB() int
//}
//
//type Work struct {
//	i int
//}
//
//func (w Work) ShowA() int {
//	return w.i + 10
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main() {
//	c := Work{3}
//	var a A = c
//	var b B = c
//	fmt.Println(a.ShowB())
//	fmt.Println(b.ShowA())
//}
//func main() {
//	i := 65
//	fmt.Println(string(i))
//}
//func main() {
//	i := -5
//	j := +5
//	fmt.Printf("%+d %+d", i, j)
//}
//func main() {
//	s := make(map[string]int)
//	delete(s, "h")
//	fmt.Println(s["h"])
//}
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person = &Person{29}
//}
//
//type B interface {
//	ShowB() int
//}
//
//type Work struct {
//	i int
//}
//
//func (w Work) ShowA() int {
//	return w.i + 10
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main() {
//	c := Work{3}
//	var a A = c
//	var b B = c
//	fmt.Println(a.ShowB())
//	fmt.Println(b.ShowA())
//}
//func main() {
//	i := 65
//	fmt.Println(string(i))
//}
//func main() {
//	i := -5
//	j := +5
//	fmt.Printf("%+d %+d", i, j)
//}
//func main() {
//	s := make(map[string]int)
//	delete(s, "h")
//	fmt.Println(s["h"])
//}
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person = &Person{29}
//}
//func incr(p *int) int {
//	*p++
//	return *p
//}
//
//func main() {
//	p :=1
//	incr(&p)
//	fmt.Println(p)
//}
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person.age = 29
//}
//func main() {
//	var a = [5]int{1, 2, 3, 4, 5}
//	var r [5]int
//
//	for i, v := range &a {
//		if i == 0 {
//			a[1] = 12
//			a[2] = 13
//		}
//		r[i] = v
//	}
//	fmt.Println("r = ", r)
//	fmt.Println("a = ", a)
//}
//func f(n int) (r int) {
//	defer func() {
//		r += n
//		recover()
//	}()
//
//	var f func()
//
//	defer f()
//	f = func() {
//		r += 2
//	}
//	return n + 1
//}
//
//func main() {
//	fmt.Println(f(3))
//}
//var p *int
//func foo() (*int, error) {
//	var i int = 5
//	return &i, nil
//}
//
//func bar() {
//	//use p
//	fmt.Println(*p)
//}
//
//func main() {
//	p, err := foo()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	bar()
//	fmt.Println(*p)
//}
//func main() {
//
//	s1 := []int{1, 2, 3}
//	s2 := s1[1:]
//	s2[1] = 4
//	fmt.Println(s1)
//	s2 = append(s2, 5, 6, 7)
//	fmt.Println(s1)
//	fmt.Println(s2)
//}
//func change(s ...int) {
//   s = append(s,3)
//}
//
//func main() {
//   slice := make([]int,5,5)
//   slice[0] = 1
//   slice[1] = 2
//   change(slice...)
//   fmt.Println(slice)
//   change(slice[0:2]...)
//   fmt.Println(slice)
//}
//func main() {
//	a := 1
//	b := 2
//	defer calc("1", a, calc("10", a, b))
//	a = 0
//	defer calc("2", a, calc("20", a, b))
//	b = 1
//}
//
//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//type People interface {
//	Speak(string) string
//}
//
//type Student struct{}
//
//func (stu *Student) Speak(think string) (talk string) {
//	if think == "speak" {
//		talk = "speak"
//	} else {
//		talk = "hi"
//	}
//	return
//}
//
//func main() {
//	var peo People = Student{}
//	think := "speak"
//	fmt.Println(peo.Speak(think))
//}

//func main() {
//	runtime.GOMAXPROCS(1)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//			time.Sleep(time.Second)
//		}
//	}()
//	time.Sleep(time.Second*10)
//}
//var a bool = true
//func main() {
//	defer func(){
//		fmt.Println("1")
//	}()
//	if a == true {
//		fmt.Println("2")
//		return
//	}
//	defer func(){
//		fmt.Println("3")
//	}()
//}
//type Foo struct {
//	bar string
//}
//func main() {
//	s1 := []Foo{
//		{"A"},
//		{"B"},
//		{"C"},
//	}
//	s2 := make([]*Foo, len(s1))
//	for i, value := range s1 {
//		s2[i] = &value
//	}
//	fmt.Println(s1[0], s1[1], s1[2])
//	fmt.Println(s2[0], s2[1], s2[2])
//}
//func main() {
//	v := []int{1, 2, 3}
//	for i := range v {
//		v = append(v, i)
//		fmt.Println(i)
//	}
//	fmt.Println(v)
//}
//var p *int
//
//func foo() (*int, error) {
//	var i int = 5
//	return &i, nil
//}
//
//func bar() {
//	//use p
//	fmt.Println(*p)
//}
//
//func main() {
//	p, err := foo()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	bar()
//	fmt.Println(*p)
//}
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main() {
//	var a A = Work{3}
//	s := a
//	fmt.Println(s.ShowA())
//	fmt.Println(s.ShowB())
//}
//type A interface {
//	ShowA() int
//}
//
//type B interface {
//	ShowB() int
//}
//
//type Work struct {
//	i int
//}
//
//func (w Work) ShowA() int {
//	return w.i + 10
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main() {
//	c := Work{3}
//	var a A = c
//	var b B = c
//	fmt.Println(a.ShowB())
//	fmt.Println(b.ShowA())
//}
//func main() {
//	i := 65
//	fmt.Println(string(i))
//}
//func main() {
//	i := -5
//	j := +5
//	fmt.Printf("%+d %+d", i, j)
//}
//func main() {
//	s := make(map[string]int)
//	delete(s, "h")
//	fmt.Println(s["h"])
//}
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person = &Person{29}
//}
//func incr(p *int) int {
//	*p++
//	return *p
//}
//
//func main() {
//	p :=1
//	incr(&p)
//	fmt.Println(p)
//}
//type Person struct {
//	age int
//}
//
//func main() {
//	person := &Person{28}
//
//	// 1.
//	defer fmt.Println(person.age)
//
//	// 2.
//	defer func(p *Person) {
//		fmt.Println(p.age)
//	}(person)
//
//	// 3.
//	defer func() {
//		fmt.Println(person.age)
//	}()
//
//	person.age = 29
//}
//func main() {
//	var a = [5]int{1, 2, 3, 4, 5}
//	var r [5]int
//
//	for i, v := range &a {
//		if i == 0 {
//			a[1] = 12
//			a[2] = 13
//		}
//		r[i] = v
//	}
//	fmt.Println("r = ", r)
//	fmt.Println("a = ", a)
//}
//func f(n int) (r int) {
//	defer func() {
//		r += n
//		recover()
//	}()
//
//	var f func()
//
//	defer f()
//	f = func() {
//		r += 2
//	}
//	return n + 1
//}
//
//func main() {
//	fmt.Println(f(3))
//}
//var p *int
//func foo() (*int, error) {
//	var i int = 5
//	return &i, nil
//}
//
//func bar() {
//	//use p
//	fmt.Println(*p)
//}
//
//func main() {
//	p, err := foo()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	bar()
//	fmt.Println(*p)
//}
//func main() {
//
//	s1 := []int{1, 2, 3}
//	s2 := s1[1:]
//	s2[1] = 4
//	fmt.Println(s1)
//	s2 = append(s2, 5, 6, 7)
//	fmt.Println(s1)
//	fmt.Println(s2)
//}
//func change(s ...int) {
//   s = append(s,3)
//}
//
//func main() {
//   slice := make([]int,5,5)
//   slice[0] = 1
//   slice[1] = 2
//   change(slice...)
//   fmt.Println(slice)
//   change(slice[0:2]...)
//   fmt.Println(slice)
//}
//func main() {
//	a := 1
//	b := 2
//	defer calc("1", a, calc("10", a, b))
//	a = 0
//	defer calc("2", a, calc("20", a, b))
//	b = 1
//}
//
//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//type People interface {
//	Speak(string) string
//}
//
//type Student struct{}
//
//func (stu *Student) Speak(think string) (talk string) {
//	if think == "speak" {
//		talk = "speak"
//	} else {
//		talk = "hi"
//	}
//	return
//}
//
//func main() {
//	var peo People = Student{}
//	think := "speak"
//	fmt.Println(peo.Speak(think))
//}

//func main() {
//	runtime.GOMAXPROCS(1)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//			time.Sleep(time.Second)
//		}
//	}()
//	time.Sleep(time.Second*10)
//}
//var a bool = true
//func main() {
//	defer func(){
//		fmt.Println("1")
//	}()
//	if a == true {
//		fmt.Println("2")
//		return
//	}
//	defer func(){
//		fmt.Println("3")
//	}()
//}
//type Foo struct {
//	bar string
//}
//func main() {
//	s1 := []Foo{
//		{"A"},
//		{"B"},
//		{"C"},
//	}
//	s2 := make([]*Foo, len(s1))
//	for i, value := range s1 {
//		s2[i] = &value
//	}
//	fmt.Println(s1[0], s1[1], s1[2])
//	fmt.Println(s2[0], s2[1], s2[2])
//}
//func main() {
//	v := []int{1, 2, 3}
//	for i := range v {
//		v = append(v, i)
//		fmt.Println(i)
//	}
//	fmt.Println(v)
//}
//func main() {
//fmt.Println("hello world")
//fmt.Println("hello world")
//}
//func main() {
//
//	var m = map[string]int{
//		"A": 21,
//		"B": 22,
//		"C": 23,
//	}
//	counter := 0
//	for k, v := range m {
//		if counter == 0 {
//			delete(m, "A")
//		}
//		counter++
//		fmt.Println(k, v)
//	}
//	fmt.Println("counter is ", counter)
//}
