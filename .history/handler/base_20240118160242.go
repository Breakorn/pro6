package handler

import (
	"net/http"
	"pro6/model"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	// logo.INFO("当前")
	var user model.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isExit := model.Login_Search(&user); !isExit {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号或密码错误"})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"success": "登陆成功！"})

	// fmt.Printf("user.UserName: %v\n", user.UserName)
	// fmt.Printf("user.Passward: %v\n", user.Passward)
	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}

func Register(ctx *gin.Context) {
	var user model.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "绑定错无"})
		return
	}
	if isExit := model.Register_Search(&user); isExit {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "该账号已注册"})
		return
	}

	if insert := model.Register_Insert(&user); insert {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": "注册成功"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "注册失败"})
}
