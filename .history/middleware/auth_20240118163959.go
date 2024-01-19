package middleware

import (
	"fmt"
	"pro6/logo"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	CurrentToken map[string]bool
}

func NewToken() *Token {
	return &Token{
		CurrentToken: make(map[string]bool),
	}
}

func (t *Token) ParseToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (t *Token) SetToken() gin.HandlerFunc {
	// token := jwt.New(jwt.SigningMethodHS256)

	// // 设置一些声明
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = "john_doe"
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 过期时间为1小时

	// // 生成签名字符串
	// tokenString, err := token.SignedString([]byte("your-secret-key"))
	// if err != nil {
	// 	fmt.Println("Error creating token:", err)
	// }

	return func(ctx *gin.Context) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		tokenString, err := token.SignedString([]byte("your-secret-key"))
		// fmt.Printf("tokenString: %v\n", tokenString)
		// fmt.Printf("tokenString: %v\n")
		// fmt.Printf("tokenString: %v\n", ctx.Request.Header["Authorization"])
		// fmt.Print("")
		fmt.Printf("ctx.Keys: %v\n", ctx.Keys)
		logo.WARN("请求前")
		username := ctx.Keys["username"]
		pwd := ctx.Keys["password"]

		ctx.Next()

		// ctx.Writer.Body
		logo.WARN("请求后")

	}
}
