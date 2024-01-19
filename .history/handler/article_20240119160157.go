package handler

import "github.com/gin-gonic/gin"

func GetArticle(ctx *gin.Context) {
	name := ctx.Keys["username"]

	ctx.JSON(200, gin.H{
		"name": name,
	})
}

func WriteAtricles(ctx *gin.Context) {
	var article *Article
	name := ctx.Keys["username"]

}
