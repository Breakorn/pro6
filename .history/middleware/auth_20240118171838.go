package middleware

import (
	"fmt"
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

	// // 要验证的 token 字符串
	// tokenString := "your-generated-jwt-token"

	// // 解析 token
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// 验证签名算法是否正确
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return []byte("your-secret-key"), nil
	// })

	// // 处理解析错误
	// if err != nil {
	// 	fmt.Println("Error parsing token:", err)
	// }

	// // 验证 token 是否有效
	// if token.Valid {
	// 	fmt.Println("Token is valid")
	// } else {
	// 	fmt.Println("Token is invalid")
	// }

	return func(ctx *gin.Context) {
		t := ctx.Request.Header["Authorization"]
		tokenString := t[0]
		fmt.Printf("tokenString: %v\n", tokenString)
		// tokenString := t.(string)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 	// 验证签名算法是否正确
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("12啦"), nil
		})
		fmt.Printf("token: %v\n", token)

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
