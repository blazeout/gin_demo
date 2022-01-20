package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello Golang",
	})
}

func main() {
	r := gin.Default()
	// 指定用户使用GET请求访问/hello时.执行sayHello这个函数
	r.GET("/hello", sayHello)
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	// 启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
}
