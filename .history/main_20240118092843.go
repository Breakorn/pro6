package main

import "github.com/dgrijalva/jwt-go"

// "pro6/db"
// "pro6/logo"
// "pro6/router"

func main() {
	// logo.NewLog()
	// db.NewDb()
	// router.Run(":8088")
	t = jwt.New(jwt.SigningMethodES256)
	s = t.SignedString(key)
}
