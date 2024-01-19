package router

import "github.com/gin-gonic/gin"

func initRouter() {

	r := gin.Default()
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
