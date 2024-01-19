package model

import (
	"fmt"
	"pro6/db"
	"pro6/logo"
	"time"
)

//	type UserInfo struct {
//		UserName string `json:"userName"`
//		PassWord string `json:"password"`
//	}
type Login struct {
	UserName string `json:"userName" binding:"required"`
	Passward string `json:"passward" binding:"required"`
}

func Register_Search(user *Login) bool {
	//查询
	name := user.UserName
	str := fmt.Sprintf("SELECT username FROM user WHERE username = %v", name)
	// str :=
	rows := db.QueryRowDB(str)
	if err := rows.Scan(&name); err != nil {
		logo.WARN(err)
		return false
	}
	return true
}

func Register_Insert(user *Login) bool {
	name := user.UserName
	pwd := user.Passward
	times := time.Now()
	count, err := db.ModifyDB("insert into users(username,password,creatime,avatar ) values (?,?,?,?)", name, pwd, times, "")
	if err != nil {
		logo.WARN(err)
		return false
	}
	logo.INFO(count)
	return true
}
