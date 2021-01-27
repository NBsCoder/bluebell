package logic

import (
	"bluelell/dao/mysql"
	"bluelell/model"
	"bluelell/pkg/snowflake"
)

//注册
func SignUp() {
	var p model.ParamSignUp
	//1.判断用户名是否存在
	mysql.QueryUserByUsername(p.Username)
	//2.生成UID
	snowflake.GenID()
	//3.用户密码加密
	//4.数据保存进数据库
	mysql.InsertUser()
}
