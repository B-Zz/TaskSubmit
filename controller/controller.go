package controller

import (
	"TaskSubmit/dao"
	"TaskSubmit/model"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册函数
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password_ := c.PostForm("password_")
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	qq := c.PostForm("qq")
	birthdate := c.PostForm("birthdate")

	user_ := dao.Manager.GetUsername(username)

	//注册验证（用户名+密码）
	switch {
	case user_.Username != "":
		c.HTML(200, "front/register.html", "用户名已存在")
	case username == "":
		c.HTML(200, "front/register.html", "请输入用户名！")
	case password == "":
		c.HTML(200, "front/register.html", "请输入密码！")
	case password_ != password:
		c.HTML(200, "front/register.html", "两次输入的密码不同！")
	default:
		{
			//使用sha256对密码加密
			password1 := sha256.Sum256([]byte(password))
			password = fmt.Sprintf("%x", password1)

			//生成cookie
			password1 = sha256.Sum256([]byte(password))
			password2 := fmt.Sprintf("%x", password1)
			cookie_ := sha256.Sum256([]byte(username + password2))
			cookie := fmt.Sprintf("%x", cookie_)

			user := model.User{
				Username:  username,
				Password:  password,
				Gender:    gender,
				Email:     email,
				Qq:        qq,
				Birthdate: birthdate,
				Cookie:    cookie,
			}
			dao.Manager.Register(&user)
			c.HTML(200, "front/home.html", "注册成功！")
		}
	}
}

// Login 登录函数
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.Manager.GetUsername(username)

	//获取密码sha256 密文
	password1 := sha256.Sum256([]byte(password))
	password = fmt.Sprintf("%x", password1)

	//登录验证
	switch {
	case user.Username == "":
		c.HTML(200, "front/login.html", "用户名不存在！")
	case user.Password != password:
		c.HTML(200, "front/login.html", "密码错误！")
	default:
		{
			// 写入cookie
			cookie, err := c.Cookie("login")
			if err != nil {
				password1 = sha256.Sum256([]byte(password))
				password2 := fmt.Sprintf("%x", password1)
				cookie = username + password2
				c.SetCookie("login", cookie, 60*60, "/", "localhost", false, true)
			}

			//验证是否为管理员
			if user.Username == "admin" {
				c.Redirect(301, "/admin")
			} else {
				c.Redirect(301, "/userinfo")
			}
		}

	}
}

// AuthMiddleWare cookie验证中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie1, err := c.Cookie("login")
		if err == nil {
			cookie2 := sha256.Sum256([]byte(cookie1))
			cookie_ := fmt.Sprintf("%x", cookie2)
			user := dao.Manager.GetUserCookie(cookie_)
			if user.Cookie == cookie_ {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败"})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "没有权限访问"})
			c.Abort()
		}
	}
}

// Search 用户信息搜索函数
func Search(c *gin.Context) {
	username := c.PostForm("username")
	user := dao.Manager.GetUsername(username)
	var info string
	if user.Username == "" {
		info = " 用户名不存在"
	}
	c.HTML(200, "back/search.html", gin.H{
		"info":      info,
		"username":  user.Username,
		"gender":    user.Gender,
		"qq":        user.Qq,
		"email":     user.Email,
		"birthdate": user.Birthdate,
	})
}
