package router

import (
	"pro6/handler"
	"pro6/logo"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	r := gin.Default()
	// 处理静态文件
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")
	// r.Use()
	// 无需token
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	return r
}

func Run(path string) {
	r := initRouter()
	logo.INFO("启动项目")
	r.Run(path)
}
