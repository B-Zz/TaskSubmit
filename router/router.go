package router

import (
	"TaskSubmit/controller"
	"TaskSubmit/dao"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	//首页
	r.GET("/", func(c *gin.Context) {

		c.HTML(200, "front/home.html", nil)
	})

	//注册界面
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "front/register.html", nil)
	})
	r.POST("/register", controller.Register)

	//登入界面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "front/login.html", nil)
	})
	r.POST("/login", controller.Login)

	//个人信息界面
	r.GET("/userinfo", controller.AuthMiddleWare(), func(c *gin.Context) {

		cookie1, _ := c.Cookie("login")
		cookie2 := sha256.Sum256([]byte(cookie1))
		cookie := fmt.Sprintf("%x", cookie2)
		user := dao.Manager.GetUserCookie(cookie)

		c.HTML(200, "front/userinfo.html", gin.H{
			"username":  user.Username,
			"gender":    user.Gender,
			"qq":        user.Qq,
			"email":     user.Email,
			"birthdate": user.Birthdate,
		})
	})
	//退出
	r.POST("/userinfo", func(c *gin.Context) {
		//删除cookie
		c.SetCookie("login", "quit", -1, "/", "localhost", false, true)
		//返回主界面
		c.Redirect(301, "/")
	})

	//后台admin组
	admin := r.Group("/admin")

	//admin主界面
	admin.GET("/", controller.AuthMiddleWare(), func(c *gin.Context) {
		c.HTML(200, "back/admin.html", nil)
	})
	//退出
	admin.POST("/", func(c *gin.Context) {
		c.SetCookie("login", "quit", -1, "/", "localhost", false, true)
		c.Redirect(301, "/")
	})

	//用户搜索界面
	admin.GET("/search", controller.AuthMiddleWare(), func(c *gin.Context) {
		c.HTML(200, "back/search.html", nil)
	})
	admin.POST("/search", controller.Search)

	r.Run()
}
