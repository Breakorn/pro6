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

	return func(ctx *gin.Context) {
		t1 := ctx.Request.Header["Authorization"]
		if len(t1) == 0 {
			ctx.JSON(200, gin.H{
				"error": "无权限",
			})

			ctx.Abort()
			return
		}
		tokenString := t1[0]
		var name string
		fmt.Printf("tokenString: %v\n", tokenString)
		// tokenString := t.(string)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 	// 验证签名算法是否正确
			fmt.Printf("token: %v\n", token)
			// fmt.Printf("token.Header[\"alg\"]: %v\n", token.Header["alg"])
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			name = t.CurrentToken[tokenString]
			return []byte(name), nil
		})
		fmt.Printf("token: %v\n", token)
		if token.Valid {
			ctx.Keys = make(map[string]any, 0)
			ctx.Keys["username"] = name
		} else {
			ctx.JSON(200, gin.H{
				"error": "无权限",
			})

			ctx.Abort()
		}

	}
}

func (t *Token) SetToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {
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
			t.CurrentToken[tokenString] = name
		} else {
			logo.ERROR("类型转换失败")
		}
		ctx.Next()
	}
}
