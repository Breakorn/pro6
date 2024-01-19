package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mbndr/logo"
)

var log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)

func main() {
	r := gin.Default()
	r.StaticFile("/", "./static/index.html")
	r.Static("/assets", "./static/assets")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	log.Info("启动项目")

	log.Debug("First debug", " and another string to log")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

}
