package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//println("已经授权")
		//context.Next()
		_,e:=context.Request.Cookie("user_cookie")
		if e==nil{
			context.Next()
		}else {
			context.Abort()
			context.HTML(http.StatusUnauthorized,"401.tmpl",nil)
		}
	}
}

