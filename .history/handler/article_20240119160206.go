package handler

import (
	"pro6/model"

	"github.com/gin-gonic/gin"
)

func GetArticle(ctx *gin.Context) {
	name := ctx.Keys["username"]

	ctx.JSON(200, gin.H{
		"name": name,
	})
}

func WriteAtricles(ctx *gin.Context) {
	var article *model.Article
	name := ctx.Keys["username"]

}
