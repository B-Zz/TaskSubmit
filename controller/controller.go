package controller

import (
	"TaskSubmit/dao"
	"TaskSubmit/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password_ := c.PostForm("password_")
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	qq := c.PostForm("qq")
	birthdate := c.PostForm("birthdate")

	user := model.User{
		Username:  username,
		Password:  password,
		Gender:    gender,
		Email:     email,
		Qq:        qq,
		Birthdate: birthdate,
	}
	user_ := dao.Manager.GetUsername(username)
	if user_.Username != "" {
		c.HTML(200, "front/register.html", "用户名已存在")
	} else {

		if username == "" {
			c.HTML(200, "front/register.html", "请输入用户名！")
		} else {
			if password == "" {
				c.HTML(200, "front/register.html", "请输入密码！")
			} else {
				if password_ != password {
					c.HTML(200, "front/register.html", "两次输入的密码不同！")
				} else {
					dao.Manager.Register(&user)
					c.HTML(200, "front/home.html", "注册成功！")
				}
			}
		}
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.Manager.GetUsername(username)

	if user.Username == "" {
		c.HTML(200, "front/login.html", "用户名不存在！")
	} else {
		if user.Password != password {
			c.HTML(200, "front/login.html", "密码错误！")
		} else {
			c.HTML(200, "front/userinfo.html", gin.H{
				"username":  user.Username,
				"gender":    user.Gender,
				"qq":        user.Qq,
				"email":     user.Email,
				"birthdate": user.Birthdate,
			})
		}
	}
}
