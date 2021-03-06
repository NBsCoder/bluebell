package logic

import (
	"bluelell/dao/mysql"
	"bluelell/model"
	"bluelell/pkg/snowflake"
)

// SignUp 注册
func SignUp(p *model.ParamSignUp) (err error) {
	//1.判断用户名是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.数据保存进数据库
	return mysql.InsertUser(user)
}

// Login 登陆
func Login(p *model.ParamLogin) error {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
