package main

import (
	"pro6/logo"
	"pro6/router"
)

func main() {
	logo.NewLog()
	router.Run(":8088")
}
