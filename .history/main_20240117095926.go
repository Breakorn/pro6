package main

import (
	"os"

	"github.com/mbndr/logo"
)

var log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)

func main() {

	log.Debug("First debug", " and another string to log")
	log.Info("Information")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

}
