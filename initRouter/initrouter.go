package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
import "net/http"

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	//添加自定义的logger中间件
	router.Use(middleware.Logger(), gin.Recovery())
	//样式文件加载

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))

	index := router.Group("/")
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
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	//添加文章
	articleRouter := router.Group("/")
	{
		// 通过获取单篇文章
		articleRouter.GET("/article/:id", handler.GetOne)
		// 获取所有文章
		articleRouter.GET("/articles", handler.GetAll)
		// 添加一篇文章
		articleRouter.POST("/article", handler.Insert)
		articleRouter.DELETE("/article/:id", handler.DeleteOne)
	}

	//seag集成
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

//func retHelloGinAndMethod(context *gin.Context) {
//	context.String(http.StatusOK, "Hello gin "+strings.ToLower(context.Request.Method)+" method")
//}
