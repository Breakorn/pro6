package model

import (
	"fmt"
	"pro6/db"
	"pro6/logo"
	"time"

	"github.com/jehiah/go-strftime"
)

//	type UserInfo struct {
//		UserName string `json:"userName"`
//		PassWord string `json:"password"`
//	}
type Login struct {
	UserName string `json:"userName" binding:"required"`
	Passward string `json:"passward" binding:"required"`
}

func Login_Search(user *Login) bool {
	//查询
	name := user.UserName
	pwd := user.Passward
	str := fmt.Sprintf("SELECT username FROM userInfo WHERE username = '%v' AND password = '%v'", name, pwd)
	// str :=
	fmt.Printf("str: %v\n", str)
	rows := db.QueryRowDB(str)
	if err := rows.Scan(&name); err != nil {
		logo.WARN(err)
		return false
	}
	return true
}

func Register_Search(user *Login) bool {
	//查询
	name := user.UserName
	str := fmt.Sprintf("SELECT username FROM userInfo WHERE username = '%v'", name)
	// str :=
	fmt.Printf("str: %v\n", str)
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
	currentTime := time.Now()

	// 使用 strftime 格式化时间
	customFormat := strftime.Format("%Y-%m-%d %H:%M:%S", currentTime)
	count, err := db.ModifyDB("insert into userInfo(username,password,creatime,avatar ) values (?,?,?,?)", name, pwd, customFormat, "")
	if err != nil {
		logo.WARN(err)
		return false
	}
	logo.INFO(count)
	return true
}
