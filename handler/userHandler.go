package handler

import "github.com/gin-gonic/gin"
import "net/http"
import "GinHello/model"

func UserSave(context *gin.Context){
	username:=context.Param("name")
	context.String(http.StatusOK,"用户"+username+"已保存")
	return
}

// 通过 query 方法进行获取参数  /user?name=daksj&age=41
func UserSaveByQuery(context *gin.Context)  {
	username:=context.Query("name")
	age:=context.Query("age")
	context.String(http.StatusOK,"用户"+username+"已保存！年纪为"+age)
}

func UserRegister(context *gin.Context)  {
	//email := context.PostForm("email")
	//password := context.PostForm("password")
	//passwordAgain := context.PostForm("password-again")

	var user model.userModel
	if err:=context.ShouldBind(&user);err!=nil{
		println("err->",err.Error())
		return
	}

	println("email", , "password", password, "password again", passwordAgain)
}
