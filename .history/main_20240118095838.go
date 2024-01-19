package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// "pro6/db"
// "pro6/logo"
// "pro6/router"

func main() {
	// logo.NewLog()
	// db.NewDb()
	// router.Run(":8088")

	// // 创建一个 token
	// token := jwt.New(jwt.SigningMethodHS256)

	// // 设置一些声明
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = "john_doe"
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 过期时间为1小时

	// // 生成签名字符串
	// tokenString, err := token.SignedString([]byte("your-secret-key"))
	// if err != nil {
	// 	fmt.Println("Error creating token:", err)
	// 	return
	// }

	// fmt.Println("JWT Token:", tokenString)

	// 验证
	// 要验证的 token 字符串
	tokenString := "your-generated-jwt-token"

	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})

	// 处理解析错误
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// 验证 token 是否有效
	if token.Valid {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is invalid")
	}
}
