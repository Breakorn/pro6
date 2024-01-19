package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// "pro6/db"
// "pro6/logo"
// "pro6/router"

func main() {
	// logo.NewLog()
	// db.NewDb()
	// router.Run(":8088")

	// 创建一个 token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置一些声明
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "john_doe"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 过期时间为1小时

	// 生成签名字符串
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("JWT Token:", tokenString)

	// 验证

}
