package middleware

import (
	"pro6/logo"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	CurrentToken map[string]string
}

func NewToken() *Token {
	return &Token{
		CurrentToken: make(map[string]string),
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

		logo.WARN("请求前")
		username := ctx.Keys["username"]
		name, flag := username.(string)
		if flag {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
			tokenString, err := token.SignedString([]byte(name))
			if err != nil {
				logo.ERROR("生成TOKEN失败")
			}
			ctx.Keys["token"] = tokenString
			t.CurrentToken[name] = name
		} else {
			logo.ERROR("类型转换失败")
		}

		ctx.Next()

		// ctx.Writer.Body
		logo.WARN("请求后")

	}
}
