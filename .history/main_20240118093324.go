package main

import (
	"fmt"
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
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
}
