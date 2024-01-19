package main

import (
	"os"

	"github.com/mbndr/logo"
)

var log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)

func main() {
	log.Info("启动项目")

	log.Debug("First debug", " and another string to log")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

}
