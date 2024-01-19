package handler

import (
	"fmt"
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

	fmt.Printf("user.UserName: %v\n", user.UserName)
	fmt.Printf("user.Passward: %v\n", user.Passward)
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

}
