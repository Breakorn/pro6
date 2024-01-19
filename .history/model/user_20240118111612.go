package model

import (
	"fmt"
	"pro6/db"
)

//	type UserInfo struct {
//		UserName string `json:"userName"`
//		PassWord string `json:"password"`
//	}
type Login struct {
	UserName string `json:"userName" binding:"required"`
	Passward string `json:"passward" binding:"required"`
}

func Login_Search(user *Login) {
	//查询
	name := user.UserName
	str := fmt.Sprintf("SELECT username FROM user WHERE username = %v", name)
	// str :=
	db.QueryRowDB(str)
}
