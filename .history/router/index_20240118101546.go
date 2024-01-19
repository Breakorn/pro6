package router

import (
	"pro6/handler"
	"pro6/logo"
	"pro6/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	// 初始化token
	token := middleware.NewToken()

	r := gin.Default()
	// 处理静态文件
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")
	// r.Use()
	// 无需token
	r.POST("/register", handler.Register)
	r.POST("/login", token.SetToken(), handler.Login)

	return r
}

func Run(path string) {
	r := initRouter()
	logo.INFO("启动项目")
	r.Run(path)
}