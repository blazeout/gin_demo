package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取form表单提交的参数

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("lesson5_form/login.html", "lesson5_form/home.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// login post请求
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		username1 := c.DefaultPostForm("username", "somebody")
		password1 := c.DefaultPostForm("xxx", "***")
		c.HTML(http.StatusOK, "home.html", gin.H{
			"Name":     username1,
			"Password": password1,
		})
	})
	r.Run(":9090")
}
