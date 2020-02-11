package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
	"strings"
)
import "net/http"

func SetupRouter() *gin.Engine{
	router:=gin.Default()

	//样式文件加载
	if mode:=gin.Mode();mode==gin.TestMode{
		router.LoadHTMLGlob("./../templates/*")
	}else{
		router.LoadHTMLGlob("templates/*")
	}
	router.Static("/statics","./statics")
	router.StaticFile("/favicon.ico","./favicon.ico")

	index:=router.Group("/")
	{
		//// 添加 Get 请求路由
		//index.GET("", retHelloGinAndMethod)
		//// 添加 Post 请求路由
		//index.POST("", retHelloGinAndMethod)
		//// 添加 Put 请求路由
		//index.PUT("", retHelloGinAndMethod)
		//// 添加 Delete 请求路由
		//index.DELETE("", retHelloGinAndMethod)
		//// 添加 Patch 请求路由
		//index.PATCH("", retHelloGinAndMethod)
		//// 添加 Head 请求路由
		//index.HEAD("", retHelloGinAndMethod)
		//// 添加 Options 请求路由
		//index.OPTIONS("", retHelloGinAndMethod)
		index.Any("", handler.Index)
	}

	// 添加 user
	userRouter:=router.Group("/user")
	{
		userRouter.GET("/:name",handler.UserSave)
		userRouter.GET("",handler.UserSaveByQuery)
		userRouter.POST("/register",handler.UserRegister)
	}

	return router
}

func retHelloGinAndMethod(context *gin.Context){
	context.String(http.StatusOK,"Hello gin "+strings.ToLower(context.Request.Method)+" method")
}