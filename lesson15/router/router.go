package router

import (
	"gin_demo/lesson15/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	// 1. 启动一个默认路由
	r := gin.Default()
	// 2. 加载前端静态文件渲染网页
	r.Static("/static", "lesson15/static")
	// 3. 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("lesson15/templates/index.html")
	r.StaticFile("/favicon.ico", "lesson15/templates/favicon.ico")
	// 4. 处理浏览器发送的GET请求
	r.GET("/bubble", controller.BubbleHandle)
	// 代办事项
	v1Group := r.Group("v1")
	{
		// 4.1 添加发送的POST请求
		v1Group.POST("/todo", controller.CreateTodo)
		// 4.2 查看所有的待办事项
		v1Group.GET("/todo", controller.LookTodoList)
		// 4.3 修改状态为完成或者未完成
		v1Group.PUT("todo/:id", controller.UpdateTodo)
		// 4.4 删除已经完成的待办事项
		v1Group.DELETE("todo/:id", controller.DeleteTodo)
	}

	return r
}
