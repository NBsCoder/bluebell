package mysql

import (
	"bluelell/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "liuhailang"

//根据用户名查重
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

//增加用户
func InsertUser(user *model.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//MD5对密码进行加密
func encryptPassword(oPassword string) string {
	//声明一个h
	h := md5.New()
	//传一个特殊字符进去
	h.Write([]byte(secret))
	//将老密码和特殊字符搞一下并转换成string类型返回
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
