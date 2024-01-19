package main

import (
	"pro6/db"
	"pro6/logo"
	"pro6/router"
)

// "pro6/db"
// "pro6/logo"
// "pro6/router"

func main() {
	logo.NewLog()
	db.NewDb()
	router.Run(":8088")

}
