package router

import (
	"os"
	"pro6/handler"

	"github.com/gin-gonic/gin"
	"github.com/mbndr/logo"
)

func initRouter() *gin.Engine {

	r := gin.Default()
	// 处理静态文件
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")

	// 无需token
	r.POST("/register", handler.Register)
	r.POST("login", handler.Login)

	return r
}

func Run(path string) {
	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
	r := initRouter()
	log.Info("启动项目")
	r.Run(path)
}
