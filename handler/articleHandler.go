package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 提交新的文章内容
// @Id 1
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.ArticleModel true "文章"
// @Success 200 object model.ResultModel 成功后返回值
// @Failure 409 object model.ResultModel 添加失败
// @Router /article [post]
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

// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 {object} model.ResultModel 成功后返回值
// @Router /article/{id} [get]
func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"result": model.ResultModel{
				Code:    http.StatusBadRequest,
				Message: "id 不是 int 类型, id 转换失败",
				Data:    e.Error(),
			},
		})
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.ArticleModel{
		Id: i,
	}

	art := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"result": model.ResultModel{
			Code:    http.StatusOK,
			Message: "查询成功",
			Data:    art,
		},
	})

}

// @Summary 获取所有的文章
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.ResultModel 成功后返回值
// @Router /articles [get]
func GetAll(context *gin.Context) {
	article := model.ArticleModel{}
	articles := article.FindAll()
	result := model.ResultModel{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    articles,
	}
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// @Summary 通过id删除指定文章类型
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200  object model.ResultModel 成功后返回值
// @Router /article/{id} [delete]
func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.ArticleModel{Id: i}
	article.DeleteOne()
}
