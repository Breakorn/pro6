package handler

import (
	"net/http"
	"pro6/logo"
	"pro6/model"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	logo.INFO("当前1")
	var user model.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isExit := model.Login_Search(&user); !isExit {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "账号或密码错误"})
		return
	}
	logo.INFO("当前2")
	ctx.Keys["username"] = &user.UserName
	ctx.Keys["password"] = &user.Password
	ctx.Next()
	ctx.JSON(http.StatusBadRequest, gin.H{"success": "登陆成功！"})
	logo.INFO("当前3")
}

func Register(ctx *gin.Context) {
	var user model.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
