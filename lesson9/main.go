package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin重定向
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status" : "ok",
		//})
		// 重定向到sogo.com
		//c.Redirect(http.StatusMovedPermanently, "https://www.sogo.com")
	})
	r.GET("/a", func(c *gin.Context) {
		// 访问/a 跳转到/b
		c.Request.URL.Path = "/b" // 将URL的路径强制变成/b
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9090")
}
