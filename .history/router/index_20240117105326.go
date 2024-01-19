package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mbndr/logo"
)

func initRouter() *gin.Engine {

	r := gin.Default()
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")

	return r
}

func Run(path string) {
	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
	r := initRouter()
	log.Info("启动项目")
	r.Run(path)
}
