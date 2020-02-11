package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)
import "net/http"
import "GinHello/model"

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已保存")
	return
}

// 通过 query 方法进行获取参数  /user?name=daksj&age=41
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户"+username+"已保存！年纪为"+age)
}

func UserRegister(context *gin.Context) {
	//email := context.PostForm("email")
	//password := context.PostForm("password")
	//passwordAgain := context.PostForm("password-again")

	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Println("err->", err.Error())
	}
	user.Save()
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()

	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	} else {
		log.Println("登录失败***************", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}
