package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件
// html页面上用到的样式文件.css .js 图片

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("lesson2_template/templates/posts/index.tmpl", "lesson2_template/templates/users/index.tmpl") // 模板解析
	r.LoadHTMLGlob("lesson2_template/templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		// HTTP 请求状态码
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		// HTTP 请求状态码
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": "liwenzhou.com",
		})
	})
	// 返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	} // 启动server
}
