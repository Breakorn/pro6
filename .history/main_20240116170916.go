package main

import (
	"fmt"
	"os"

	"github.com/mbndr/logo"
)

func main() {

	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
	file, err := os.Stat("./db/test.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file: %v\n", file)
	log.Debug("First debug", " and another string to log")
	log.Info("Information")
	log.Warn("Warning", " message")
	log.Error("Error message")
	log.Fatal("Fatal error", " because of something")

	// log.Debugf("Debug value %d", 16)
	// log.Infof("Listening on port %d", 8080)

}
