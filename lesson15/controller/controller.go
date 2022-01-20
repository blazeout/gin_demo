package controller

import (
	"gin_demo/lesson15/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BubbleHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项, 点击提交就会发送POST请求到这里
	// 1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	if err := models.CreateData(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// 3.返回响应
		c.JSON(http.StatusOK, todo)
	}
}

func LookTodoList(c *gin.Context) {
	// 查询todo这个表里面的所有数据
	var todoList []models.Todo
	if err := models.ShowData(&todoList); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	// 拿到id之后更新status状态即可
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := models.UpdateTodo(&todo, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	c.BindJSON(&todo)
	models.Save(todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	if err := models.DeleteTodo(&todo, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
