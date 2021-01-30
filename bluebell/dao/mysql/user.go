package mysql

import (
	"bluelell/model"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "liuhailang"

// CheckUserExist 根据用户名查重
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该用户已存在")
	}
	return
}

// InsertUser 增加用户
func InsertUser(user *model.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword MD5对密码进行加密
func encryptPassword(oPassword string) string {
	//声明一个h
	h := md5.New()
	//传一个特殊字符进去
	h.Write([]byte(secret))
	//将老密码和特殊字符搞一下并转换成string类型返回
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login 用户登陆
func Login(user *model.User) (err error) {
	oPassword := user.Password //将用户传进来的密码先存起来，待会要做比较
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username) //直接从数据库中把这个数据查找出来存到user里边
	//判断用户是否存在
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	//将传进来的密码加密，在和数据库中的密码对比
	Password := encryptPassword(oPassword)
	if Password != user.Password {
		return errors.New("密码错误")
	}
	return
}
