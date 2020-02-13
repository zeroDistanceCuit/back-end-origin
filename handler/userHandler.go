package handler

import (
	"GinHello/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
import "net/http"
import "GinHello/model"

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已保存")
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
		context.SetCookie("user_cookie", string(u.Id), 1000, "/", "localhost", false, true)
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	} else {
		log.Println("登录失败*********")
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}

func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)

	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
	}

	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误", err.Error())
	}

	file, e := context.FormFile("avatar-file")
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	e = os.MkdirAll(path, os.ModePerm)

	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = context.SaveUploadedFile(file, path+fileName)

	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	//TODO 更新服务器ip
	avatarUrl := "http://localhost:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}
