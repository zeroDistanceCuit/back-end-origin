package model

type userModel struct {
	Email         string `form:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
}
