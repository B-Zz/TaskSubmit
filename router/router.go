package router

import (
	"TaskSubmit/controller"
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

	r.Run()
}
