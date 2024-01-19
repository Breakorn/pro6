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

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(200, gin.H{
			"Message": err,
		})
	}
	name := ctx.Keys["username"]
	strName, err := name.(string)
	if !err {
		ctx.JSON(200, gin.H{
			"Message": err,
		})
	}
	isSuccess := model.Article_Insert(strName, article)
	// if isSuccess {
	ctx.JSON(200, gin.H{
		"Success": isSuccess,
	})
	// }
}
