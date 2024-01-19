package router

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {

	r := gin.Default()
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")

	return r
}

func Run(path string) {
	r := initRouter()
	r.Run(path)
}
