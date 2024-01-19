package middleware

import "github.com/gin-gonic/gin"

type Token struct {
	CurrentToken map[string]bool
}

func NewToken() *Token {
	return &Token{
		CurrentToken: make(map[string]bool),
	}
}

func (t *Token) ParseToken() {
	return func(ctx *gin.Context) {

	}
}

func (t *Token) SetToken() {

}
