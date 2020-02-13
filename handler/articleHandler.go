package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Insert(context *gin.Context) {
	article := model.ArticleModel{}
	var id = -1
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
