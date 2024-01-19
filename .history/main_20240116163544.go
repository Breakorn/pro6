package main

import (
	"os"

	"github.com/mbndr/logo"
)

func main() {
	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "prefix ", true)
	log.Debug("First debug", " and another string to log")
	log.Info("Information")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

	log.Debugf("Debug value %d", 16)
	log.Infof("Listening on port %d", 8080)
	log.Warnf("Invalid user %s", "踩踩踩")
	log.Errorf("Couldn't load config file: %s", "path")
	log.Fatalf("Fatal error: %s", err.Error())
}
