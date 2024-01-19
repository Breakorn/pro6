package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// "pro6/db"
// "pro6/logo"
// "pro6/router"

func main() {
	// logo.NewLog()
	// db.NewDb()
	// router.Run(":8088")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// tokenString, err := token.SignedString(hmacSampleSecret)

	// fmt.Println(tokenString, err)
}
